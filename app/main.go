package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kanaru-ssk/sqlboiler-sample/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@db:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Error:", err)
	}
	boil.SetDB(db)
	ctx := context.Background()

	// データを全削除
	models.Users().DeleteAll(ctx, db)
	models.Teams().DeleteAll(ctx, db)
	models.TeamMembers().DeleteAll(ctx, db)

	// ユーザーを作成
	newUsers := [10]models.User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}, {ID: 3, Name: "Charlie"}, {ID: 4, Name: "Dave"}, {ID: 5, Name: "Eve"}, {ID: 6, Name: "Frank"}, {ID: 7, Name: "Grace"}, {ID: 8, Name: "Hank"}, {ID: 9, Name: "Ivy"}, {ID: 10, Name: "Jack"}}
	for _, user := range newUsers {
		err := user.Insert(ctx, db, boil.Infer())
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// チームを作成
	newTeams := [3]models.Team{{ID: 1, Name: "Team1"}, {ID: 2, Name: "Team2"}, {ID: 3, Name: "Team3"}}
	for _, team := range newTeams {
		err := team.Insert(ctx, db, boil.Infer())
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// チームメンバーを作成
	newTeamMembers := [10]models.TeamMember{{TeamID: null.NewInt(1, true), UserID: null.NewInt(1, true), UserRole: "OWNER"}, {TeamID: null.NewInt(1, true), UserID: null.NewInt(2, true), UserRole: "MEMBER"}, {TeamID: null.NewInt(1, true), UserID: null.NewInt(3, true), UserRole: "MEMBER"}, {TeamID: null.NewInt(2, true), UserID: null.NewInt(4, true), UserRole: "OWNER"}, {TeamID: null.NewInt(2, true), UserID: null.NewInt(5, true), UserRole: "MEMBER"}, {TeamID: null.NewInt(2, true), UserID: null.NewInt(6, true), UserRole: "MEMBER"}, {TeamID: null.NewInt(3, true), UserID: null.NewInt(7, true), UserRole: "OWNER"}, {TeamID: null.NewInt(3, true), UserID: null.NewInt(8, true), UserRole: "MEMBER"}, {TeamID: null.NewInt(3, true), UserID: null.NewInt(9, true), UserRole: "MEMBER"}, {TeamID: null.NewInt(3, true), UserID: null.NewInt(10, true), UserRole: "MEMBER"}}
	for _, teamMember := range newTeamMembers {
		err := teamMember.Insert(ctx, db, boil.Infer())
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// すべてのユーザーを取得
	fmt.Println("All Users")
	users, err := models.Users().All(ctx, db)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, user := range users {
		fmt.Println("ID:", user.ID, "Name:", user.Name)
	}

	// すべてのチームを取得
	fmt.Println("All Teams")
	teams, err := models.Teams().All(ctx, db)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, team := range teams {
		fmt.Println("ID:", team.ID, "Name:", team.Name)
	}

	// すべてのチームメンバーを取得
	fmt.Println("All TeamMembers")
	teamMembers, err := models.TeamMembers().All(ctx, db)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, teamMember := range teamMembers {
		fmt.Println("TeamID:", teamMember.TeamID, "UserID:", teamMember.UserID, "UserRole:", teamMember.UserRole)
	}

	// チーム1のメンバーを取得
	fmt.Println("Team1 Members")
	team1Members, err := models.TeamMembers(models.TeamMemberWhere.TeamID.EQ(null.NewInt(1, true))).All(ctx, db)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, teamMember := range team1Members {
		fmt.Println("TeamID:", teamMember.TeamID, "UserID:", teamMember.UserID, "UserRole:", teamMember.UserRole)
	}

	// チーム1のメンバーを取得(usersテーブルと結合)
	type member struct {
		UserID   int    `boil:"user_id"`
		UserName string `boil:"user_name"`
		UserRole string `boil:"user_role"`
	}
	var teamMemberWithUser []member
	models.NewQuery(
		qm.Select("users.id as user_id", "users.name as user_name", "team_member.user_role as user_role"),
		qm.From("team_member"),
		qm.InnerJoin("users on users.id = team_member.user_id "),
		qm.Where("team_member.team_id = ?", 1),
	).Bind(ctx, db, &teamMemberWithUser)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, member := range teamMemberWithUser {
		fmt.Println("UserID:", member.UserID, ", Name:", member.UserName, ", UserRole:", member.UserRole)
	}
}
