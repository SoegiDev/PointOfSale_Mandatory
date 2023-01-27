package seeds

import (
	"PointOfSale/database"
	"PointOfSale/schema"
)

func InsertGroupMenu() {
	var total int64
	database.Database.Model(&schema.GroupMenu{}).Count(&total)
	if total == 0 {
		var data = []schema.GroupMenu{
			{Name: "Burger", Description: "Burger Group"}, {Name: "Pizza", Description: "Pizza Group"}, {Name: "Taco", Description: "Taco Group"},
			{Name: "HotDog", Description: "HotDog Group"}, {Name: "Chicken", Description: "Chicken Group"}, {Name: "Desserts", Description: "Desserts Group"}}
		database.Database.Create(&data)
	}
}

func InsertMenu() {
	var total int64
	database.Database.Model(&schema.Menu{}).Count(&total)
	if total == 0 {
		var data = []schema.Menu{
			{Name: "Beef", GroupId: 1, Price: 50000, Picture: "https://images.unsplash.com/photo-1603064752734-4c48eff53d05?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=388&q=80"},
			{Name: "Cheese", GroupId: 1, Price: 50000, Picture: "https://images.unsplash.com/photo-1603064752734-4c48eff53d05?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=388&q=80"},
			{Name: "Tuna", GroupId: 1, Price: 50000, Picture: "https://images.unsplash.com/photo-1603064752734-4c48eff53d05?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=388&q=80"},
			{Name: "Met Lover", GroupId: 2, Price: 50000, Picture: "https://images.unsplash.com/photo-1590947132387-155cc02f3212?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"},
			{Name: "Chicken Happy", GroupId: 2, Price: 50000, Picture: "https://images.unsplash.com/photo-1590947132387-155cc02f3212?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"},
			{Name: "Crispy Tuna", GroupId: 2, Price: 50000, Picture: "https://images.unsplash.com/photo-1590947132387-155cc02f3212?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"},
			{Name: "Nacked Chicken", GroupId: 3, Price: 50000, Picture: "https://images.unsplash.com/photo-1574782091246-c65ed4510300?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=465&q=80"},
			{Name: "Crunch Supreme", GroupId: 3, Price: 50000, Picture: "https://images.unsplash.com/photo-1574782091246-c65ed4510300?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=465&q=80"},
			{Name: "Grilled Stuft", GroupId: 3, Price: 50000, Picture: "https://images.unsplash.com/photo-1574782091246-c65ed4510300?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=465&q=80"},
			{Name: "Reguler", GroupId: 4, Price: 50000, Picture: "https://images.unsplash.com/photo-1524237629218-0a109277890c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80"},
			{Name: "Jumbo", GroupId: 4, Price: 50000, Picture: "https://images.unsplash.com/photo-1524237629218-0a109277890c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80"},
			{Name: "Giant", GroupId: 4, Price: 50000, Picture: "https://images.unsplash.com/photo-1524237629218-0a109277890c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80"},
			{Name: "Original", GroupId: 5, Price: 50000, Picture: "https://images.unsplash.com/photo-1626645738196-c2a7c87a8f58?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"},
			{Name: "Bulgogi", GroupId: 5, Price: 50000, Picture: "https://images.unsplash.com/photo-1626645738196-c2a7c87a8f58?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"},
			{Name: "Crispy", GroupId: 5, Price: 50000, Picture: "https://images.unsplash.com/photo-1626645738196-c2a7c87a8f58?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"},
			{Name: "Tiramisu", GroupId: 6, Price: 50000, Picture: "https://images.unsplash.com/photo-1551024506-0bccd828d307?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=464&q=80"},
			{Name: "Choco Cheese", GroupId: 6, Price: 50000, Picture: "https://images.unsplash.com/photo-1551024506-0bccd828d307?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=464&q=80"},
			{Name: "Oreo Choco", GroupId: 6, Price: 50000, Picture: "https://images.unsplash.com/photo-1551024506-0bccd828d307?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=464&q=80"}}
		database.Database.Create(&data)
	}
}

func InsertTable() {
	var total int64
	database.Database.Model(&schema.TableTransaction{}).Count(&total)
	if total == 0 {
		var data = []schema.TableTransaction{
			{Name: "Meja001", Description: "Meja 001"}, {Name: "Meja002", Description: "Meja 002"}, {Name: "Meja003", Description: "Meja 003"},
			{Name: "Meja004", Description: "Meja 004"}, {Name: "Meja005", Description: "Meja 001"}, {Name: "Meja006", Description: "Meja 006"}}
		database.Database.Create(&data)
	}
}
