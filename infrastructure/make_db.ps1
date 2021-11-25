$query = Get-Content .\infrastructure\todo-app.sql
$db = ".\infrastructure\base.sqlite3"
Invoke-SqliteQuery -Query $Query -DataSource $Database