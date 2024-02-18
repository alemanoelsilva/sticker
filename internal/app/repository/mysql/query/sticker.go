package query

const (
	AddStickerQuery        = "INSERT INTO stickers (name, description, category, frequency, status, is_public, is_auto_approved) VALUES (:name, :description, :category, :frequency, :status, :is_public, :is_auto_approved)"
	GetStickerByIdQuery    = "SELECT * FROM stickers WHERE id = '%d' LIMIT 1"
	GetStickersQuery       = "SELECT * FROM stickers"
	DeleteStickerByIdQuery = "DELETE FROM stickers where id = '%d'"
	UpdateStickerByIdQuery = "UPDATE SET name =:name, description =:description, category =:category, frequency =:frequency, status =:status, is_public =:is_public, is_auto_approved =:is_auto_approved WHERE id = '%d'"
)
