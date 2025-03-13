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
                startDate DATE,
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


# <UPDATE> employee
def update_employee(employee_id, name, age, position, start_date, sindicate, salary):
    c.execute('''
        UPDATE employees 
        SET name = ?, age = ?, position = ?, startDate = ?, sindicate = ?, salary = ? 
        WHERE employee_id = ?''', 
        (name, age, position, start_date, sindicate, salary, employee_id))
    conn.commit()

# <DELETE> employee
def delete_employee(employee_id):
    c.execute('''
                DELETE FROM employees WHERE employee_id=?''',
                (employee_id,))
    c.commit()


def menu():
    while True:
        print("\nMenu:")
        print("1. Create Employee")
        print("2. Read Employees")
        print("3. Update Employee")
        print("4. Delete Employee")
        print("5. Exit")

        choice = int(input("Enter your choice: "))

        if choice == 1:
            print('Insert employee information')
            employee_id = int(input("ID: ")) 
            name = str(input("Name: "))
            age = int(input("Age: "))
            position = str(input("Position: ")) 
            start_date = (input("Date(MM/DD/YY): ")) 
            sindicate = bool(input("Sindicate(True/False): "))
            salary = int(input("Salary: "))
        elif choice == 2:
