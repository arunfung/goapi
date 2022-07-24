package cmd

import (
	"goapi/database/migrations"
	"goapi/pkg/migrate"

	"github.com/spf13/cobra"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	// 所有 migrate 下的子命令都会执行以下代码
}

var MigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func init() {
	Migrate.AddCommand(
		MigrateUp,
		MigrateRollback,
		MigrateReset,
		MigrateRefresh,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

var MigrateRollback = &cobra.Command{
	Use: "down",
	// 设置别名 migrate down == migrate rollback
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

var MigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

var MigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}
