# endpoint

* post /products — Создать продукт / {"title": "Молоко", "price":200.0 } / sucess

* get /products — Получить все продукты / нет / [ { "id": 1, "title": "Молоко","price": 200.0 }, {"id": 2,"title": "Хлеб", "price": 50.0} ]

* get /products/{id} — Получить продукт по ID / нет / {"id": 1,"title": "Молоко","price": 200.0}

* patch /products/{id} — Изменить продукт / {"title": "Молоко2","price": 22.0 } / {"id": 1,"title": "Молоко2","price": 22.0}

* delete /products/{id} — Удалить продукт / нет / нет