import sqlite3

# Connect to SQLite database (or create it if it doesn't exist)
conn = sqlite3.connect('example.db')
c = conn.cursor()

# Create a table
c.execute('''CREATE TABLE IF NOT EXISTS items
             (id INTEGER PRIMARY KEY, name TEXT, quantity INTEGER)''')

# Function to create a new item
def create_item(name, quantity):
    c.execute("INSERT INTO items (name, quantity) VALUES (?, ?)", (name, quantity))
    conn.commit()

# Function to read all items
def read_items():
    c.execute("SELECT * FROM items")
    return c.fetchall()

# Function to update an item
def update_item(item_id, name, quantity):
    c.execute("UPDATE items SET name = ?, quantity = ? WHERE id = ?", (name, quantity, item_id))
    conn.commit()

# Function to delete an item
def delete_item(item_id):
    c.execute("DELETE FROM items WHERE id = ?", (item_id,))
    conn.commit()

# Function to display the menu and handle user input
def menu():
    while True:
        print("\nMenu:")
        print("1. Create item")
        print("2. Read items")
        print("3. Update item")
        print("4. Delete item")
        print("5. Exit")
        choice = input("Enter your choice: ")

        if choice == '1':
            name = input("Enter item name: ")
            quantity = int(input("Enter item quantity: "))
            create_item(name, quantity)
        elif choice == '2':
            items = read_items()
            for item in items:
                print(item)
        elif choice == '3':
            item_id = int(input("Enter item ID to update: "))
            name = input("Enter new item name: ")
            quantity = int(input("Enter new item quantity: "))
            update_item(item_id, name, quantity)
        elif choice == '4':
            item_id = int(input("Enter item ID to delete: "))
            delete_item(item_id)
        elif choice == '5':
            break
        else:
            print("Invalid choice. Please try again.")

# Run the menu
menu()

# Close the connection
conn.close()
