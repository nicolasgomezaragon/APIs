import pandas as pds

def read_excel(file_path):
    df =  pds.read_excel(file_path)
    return df

file_path = 'path_to_file'
data = read_excel(file_path)
print(data.head())

