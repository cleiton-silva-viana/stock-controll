# Specification of stock controll

## Features

### All Employees 
- [] Registrer new employee
    - The system to have: name, age, gender, email and cell number
    - The password to have chars, numbers and special characters
    - The password to have than 8 characters
    - The password to have maximum of de 24 characters
    - Cell phone and email to have unique in system
- [] Employee login
    - The employee to have provide email and password
    - Maximun of 3 attemps
- [] Recover account
- [] Alter employee datas
    - [] Logs to register employee update actions
- [] History of actions in system

### Only by resouce humans
- [] Delete employees account
    - If employee data for disabled than more 30 days
    - [] Logs to register employee deletion
- [] Get employee datas
    - To have return: name, registration, email and cell number
- [] Change employee position
    - [] Logs to register employee change position

### Only by sales employee
- [] Decrement product quantity
    - Quantity no fewer than 0
- [] Warn product with expiration data upcoming
    - Warn to have contain: name, quantity, date expiration and date of manufacture
- [] Warn product with breakdown
    - Warn to have contain: name, quantity

### Only by buyer employee
- [] Register new manufacturer
    - [] Logs to register datas of manufacturer created
- [] Alter manufacturer datas
    - [] Logs to register datas updates manufacturer 
- [] Register new product
    - [] Logs to register datas of products created 
- [] Remove product
    - [] Logs to register datas of products removed
- [] Disable product
    - [] Logs to register datas of products disabled
- [] Alter product datas
    - [] Logs to register datas of products updated

### Only by lecturer
- [] Receive product 
    - [] Logs to register products received
- [] Send product
    - [] Logs to register products sends
- [] Deny product
    - [] Logs to register products denied

### Only by stocklist
- [] Move product in stock
    - [] Logs to register stock movimentation
- [] Discard invalid product
    - [] Logs to register actions of discard
- [] Recicle invalid product
    - [] Logs to register actions of recycle
- [] Swap invalid product
    - [] Logs to register actions of swap

### System
- [] Routine for check due date products
    - [] Logs for register products with expiration date upcoming
- [] Create alert for product wifth expiration date upcoming
    - Send alert for employee with relationship to product movimentation and administration
- [] Routine for check products with low volumes
    - [] Logs for register products with low volumes
- [] Create tasks for priority order for execution
    - The tasks to have scope employee
- [] Create a resume for stock
