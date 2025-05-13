E_commerce (go + postgres)

Tables : 
1. user_details
2. products 
3. orders

Implementation : 
1. Transaction query for
    a. user ordering a product. (order)
    b. decrementing the stock of that product by 1.
2. Trigger for update_at col. when user update a product

main.go file run cmd : go run .
trigger.sql file run cmd : psql -h localhost -p 5433 -U postgres -d e_commerce -f trigger.sql