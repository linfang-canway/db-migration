package migrations

import (
	"gorm.io/gorm"

	"db-migration/migrator"
)

func init() {
	// add current migration to migrator
	migrator.GetMigrator().AddMigration(&migrator.Migration{
		Version: "20241110203050",
		Name:    "20241110203050_add_group",
		Mode:    migrator.GormMode,
		Up:      mig20231009103050Up,
		Down:    mig20231009103050Down,
	})
}

func mig20231009103050Up(tx *gorm.DB) error {

	type Groups struct {
		ID   uint   `gorm:"type:bigint(1) unsigned not null;primaryKey;autoIncrement:false"`
		Name string `gorm:"type:varchar(255) not null;"`
		gorm.Model
	}

	if err := tx.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").
		AutoMigrate(&Groups{}); err != nil {
		return err
	}

	return nil
}

func mig20231009103050Down(tx *gorm.DB) error {

	type Groups struct {
		ID   uint   `gorm:"type:bigint(1) unsigned not null;primaryKey;autoIncrement:false"`
		Name string `gorm:"type:varchar(255) not null;"`
		gorm.Model
	}

	if err := tx.Migrator().DropTable(Groups{}); err != nil {
		return err
	}

	return nil
}
