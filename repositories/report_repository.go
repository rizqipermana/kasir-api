package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetReportHariIni() (*models.ReportHariIni, error) {

	var report models.ReportHariIni

	// total revenue hari ini
	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&report.TotalRevenue)

	if err != nil {
		return nil, err
	}

	// total transaksi hari ini
	err = r.db.QueryRow(`
		SELECT COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&report.TotalTransaksi)

	if err != nil {
		return nil, err
	}

	// produk terlaris hari ini
	err = r.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) AS total_qty
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`).Scan(
		&report.ProdukTerlaris.Nama,
		&report.ProdukTerlaris.QtyTerjual,
	)

	// ⚠️ kalau hari ini belum ada transaksi sama sekali
	if err == sql.ErrNoRows {
		report.ProdukTerlaris.Nama = ""
		report.ProdukTerlaris.QtyTerjual = 0
		return &report, nil
	}

	if err != nil {
		return nil, err
	}

	return &report, nil
}
