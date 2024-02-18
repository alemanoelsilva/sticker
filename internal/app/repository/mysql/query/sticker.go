package query

const (
	AddStickerQuery        = "INSERT INTO stickers (name, description, category, frequency, status, is_public, is_auto_approval, user_id) VALUES (:name, :description, :category, :frequency, :status, :is_public, :is_auto_approval, :user_id)"
	GetStickerByIdQuery    = "SELECT id, name, description, category, frequency, status, is_public, is_auto_approval, user_id FROM stickers WHERE user_id = '%d' AND id = '%d' LIMIT 1"
	GetStickersQuery       = "SELECT id, name, description, category, frequency, status, is_public, is_auto_approval, user_id FROM stickers where user_id = '%d'"
	DeleteStickerByIdQuery = "DELETE FROM stickers where id = '%d'"
	UpdateStickerByIdQuery = "UPDATE SET name = :name, description = :description, category = :category, frequency = :frequency, status = :status, is_public = :is_public, is_auto_approval = :is_auto_approval WHERE id = '%d'"
)
