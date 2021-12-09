# Day1

> Learn with real example

## 1. Write a function that can analyze this line

> "ID","Subject","Status","Estimation","StartDate","DueDate"

### Expected:

1. Receive input from CLI
2. The output should have: number of columns, display a table with columns name

```md
| Column 1 | Column 2 | ... |
```

## 2. Write a function that can analyze this line

> 1,"Task 1","Done","5","2021-12-01","2021-12-01"

### Expected:

1. Receive input from CLI
2. Store receive input in a structure data type
3. The output of this function must be mapped with the 1st function above

## 3. Expand the function that can store all input lines in a structure data array

```m
1,"Task 1","Done","5","2021-12-01","2021-12-01"
2,"Task 2","Done","4","2021-12-01","2021-12-01"
3,"Task 3","Done","3","2021-12-01","2021-12-01"
4,"Task 4","Rejected","2","2021-12-01","2021-12-01"
```

### Expected:

1. Output a complete table with column names, and values for each rows
2. Generate statistic dashboard in table format for:

- Number of task by date
- Number of task by status
- Number of task by date + status
- Number of task by month + status
- Number of task by year + status

## 4. Extends the 3rd requirement

### Expected:

1. Read data from a specific file and store in application's data folder
2. User can select a file to be convert via cli command with param
3. The application should allow user to select the type of statistic to be display
4. The application should allow user to export the selected statistic to a destination file/folder
5. The application should allow user add new row to dataset by manually input a new row
6. The application should allow user add new new rows to dataset by import a new collection of data

## 5. Connect DB

### Expected:

1. Store the application's data in mysqli

## 6. UI

### Expected:

1. Convert the current application into Web form
2. Package your application with Docker
