Building a service system for retailer to sell products using the APIs. 

The following are the requirements :

1) Retailer should be able to add and update the product details
2) Customers should be able to view their products, place an order and view the history of transactions .
3) Cool down period of 5 min should be set after every transaction of the customer. If a customer places and order in the mean while ,the order status should be "processing order".

Tools and Softwares used : MySQL Workbench, Goland , Postman

Three databases namely Customer,Order and Product are created .
Customer Database contains the following fields :
1. Customer ID
2. Name
3. Email
4. Phone
5. Address

Order Database contains the following fields :
1. OrderID
2. ProductID
3. Customer ID
4. Product(Foreign key)
5. Customer(Foreign key)
6. Quantity 
7. Order Status

Product Database contains the following fields :
1. ProductID
2. Name
3. Description
4. Price
5. Quantity 

Now we create APIs for these databases

Customer:
1. GET : To get all the users
2. POST : To create new customer database
3. GET/ID : To get database of given user

Product:
1. GET : To get all the products
2. POST : To create new product database
3. GET/ID : To get database of given product
4. PATCH/ID : To update the database of given product
5. DELETE/ID : To delete the data of given product. 

Mutex is used while "GET","DELETE" and "PATCH" inorder to prevent any otber function taking place simultaneously

Order:
1. GET : To get all the order
2. POST : To create new order database
//Whenever a new order is placed , the customer is cooldown for 5 min and will not be able to complete placing another order in meanwhile
3. GET/ID : To get database of given order

Mutex is used while "GET" and "POST" inorder to prevent any otber function taking place simultaneously

Requests are handled in model files.

Then routes are setup using gin engine .

Server is set up in the main function.

Main_test file is written to perform the testing operations.

