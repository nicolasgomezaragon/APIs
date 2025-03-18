# Objetive here is to grab the extracted data,
# create the linear model with Scikit-learn
# 

# Import necessary libraries
import numpy as np
import pandas as pd
from sklearn.linear_model import LinearRegression
import matplotlib.pyplot as plt

# Step 1: Load your data (example data)
# Assuming we have historical stock prices with 'Date' and 'Price'
data = {
    "Date": ["2023-01-01", "2023-01-02", "2023-01-03", "2023-01-04", "2023-01-05"],
    "Price": [100, 102, 101, 104, 107],
}
df = pd.DataFrame(data)

# Convert 'Date' to datetime and calculate 'Days since start'
df['Date'] = pd.to_datetime(df['Date'])
df['Days'] = (df['Date'] - df['Date'].min()).dt.days

# Step 2: Prepare data for the model
X = df[['Days']]  # Independent variable (time in days)
y = df['Price']   # Dependent variable (stock price)

# Step 3: Create and train the linear regression model
model = LinearRegression()
model.fit(X, y)

# Step 4: Use the model to make predictions
df['Predicted Price'] = model.predict(X)

# Step 5: Visualize the results
plt.figure(figsize=(8, 5))
plt.scatter(df['Days'], df['Price'], color='blue', label='Actual Prices')
plt.plot(df['Days'], df['Predicted Price'], color='red', label='Linear Regression Line')
plt.xlabel('Days since Start')
plt.ylabel('Stock Price')
plt.title('Stock Price vs. Time')
plt.legend()
plt.show()

# Step 6: Evaluate the model (optional)
r_squared = model.score(X, y)
print(f"R-squared: {r_squared:.2f}")


df = pd.DataFrame(data)

#Example 0 to grab the values and convert it to 
data = {
    "Date": ["2023-01-01", "2023-01-02", "2023-01-03", "2023-01-04", "2023-01-05"],
    "Price": [100, 102, 101, 104, 107],
    "Predicted Price": [100.4, 101.6, 102.8, 104.0, 105.2],
}

# Step 1: Create the Excel file
excel_filename = "stock_prediction.xlsx"
df.to_excel(excel_filename, index=False)  # Set index=False to exclude the row numbers

# Step 2: Confirm the file is saved
print(f"Data successfully exported to {excel_filename}")