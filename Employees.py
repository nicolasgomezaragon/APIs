import sqlite3

# Establishing connection
conn = sqlite3.connect('payroll.db')
c = conn.cursor()

# Initializating table
c.execute(''' CREATE TABLE IF NOT EXISTS
           employees 
                (
                employee_id INTEGER PRIMARY KEY, 
                name TEXT, 
                age INTEGER,
                position TEXT, 
                start_date DATE,
                sindicate BOOLEAN,
                salary INTEGER
                )''')


# <CREATE> a new employee
def add_employee(employee_id, name, age, position, start_date, sindicate, salary):
    c.execute('''
              INSERT INTO employees
              (employee_id, name, age, position, start_date, sindicate, salary) 
              VALUES (?,?,?,?,?,?,?)''',
              (employee_id, name, age, position, start_date, sindicate, salary))    
    conn.commit()

# <READ> all records
def read_employees():
    c.execute("SELECT * FROM employees")
    return c.fetchall()

#######################
def read_single_employee(employee_id):
    c.execute("SELECT * FROM employees WHERE employee_id=?",(employee_id))
    return c.fetchone()
#####################


# <UPDATE> employee
def update_employee(employee_id, name, age, position, start_date, sindicate, salary):
    c.execute('''
        UPDATE employees 
        SET name = ?, age = ?, position = ?, start_date = ?, sindicate = ?, salary = ? 
        WHERE employee_id = ?''', 
        (name, age, position, start_date, sindicate, salary, employee_id))
    conn.commit()

# <DELETE> employee
def delete_employee(employee_id):
    c.execute('''
                DELETE FROM employees WHERE employee_id=?''',
                (employee_id,))
    conn.commit()


def menu():
    while True:
        print("\nMenu:")
        print("1. Create Employee")
        print("2. Read Employees")
        print("3. Update Employee")
        print("4. Delete Employee")
        print("5. Exit")

        choice = int(input("Enter your choice: "))

        # CREATE 
        if choice == 1:
            print('Insert employee information')
            employee_id = int(input("ID: ")) 
            name = str(input("Name: "))
            age = int(input("Age: "))
            position = str(input("Position: ")) 
            start_date = (input("Date(MM/DD/YY): ")) 
            sindicate = bool(input("Sindicate(True/False): "))
            salary = int(input("Salary: "))
            add_employee(employee_id, name, age, position, start_date, sindicate, salary)
        # READ
        elif choice == 2:
            employees = read_employees()
            for emp in employees:
                print(emp)
        #UPDATE
        elif choice == 3:
            employee_id = int(input("Enter ID to update data: "))

            employee = read_single_employee(employee_id)
            if employee:
                print(f"Current data: {employee}")
                name = input("Name: ") or employee[1]
                age = input("Age:") or employee[2]
                position = input("Position:") or employee[3]
                start_date = input(":") or employee[4]
                sindicate = input(":") or employee[5] 
                salary = input(":") or employee[6]
                update_employee(employee_id, name, age, position, start_date, sindicate, salary)
            else:
                print('Employe not found')
        #DELETE
        elif choice == 4:
            employee_id = int(input("Enter ID to delete employee: "))
            delete_employee(employee_id)
            print(f"Employee with ID {employee_id} has been successfully deleted")
        elif choice == 5:
            break
        else: 
            print('Invalid choice, please try again')

menu()

conn.close()

        
