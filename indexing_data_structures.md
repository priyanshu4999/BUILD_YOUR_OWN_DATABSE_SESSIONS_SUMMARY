#CHAPTER 2 INDEXING DATA STRUCTURES 
## types of queries 
> queries can be briken down into three types  
1. scan whole dataset (NO INDEX)
2. Point query , query the index by a specific key
3. Range query , query index by a range (index is sorted)





---

### 2.1 Types of SQL Queries

#### 1. Full Table Scan
**Description**: Scans the entire dataset without using any index.
- **Use Case**: Necessary when no suitable index exists or when the query needs to process every row.
- **Performance**: Generally slower for large datasets since it reads all rows.

**Example**:
```sql
SELECT * FROM Employees;
```

---

#### 2. Point Query
**Description**: Queries the index by a specific key to retrieve a specific row or set of rows.
- **Use Case**: Efficient for fetching data based on a unique key, such as an ID.
- **Performance**: Very fast for indexed columns as it quickly locates the row(s) without scanning the entire table.

**Example**:
```sql
SELECT * FROM Employees WHERE EmployeeID = 123;
```

---

#### 3. Range Query
**Description**: Queries the index by a range of values using a sorted index.
- **Use Case**: Useful for retrieving data within a specified range, such as date ranges or numeric ranges.
- **Performance**: Efficient for sorted indexes, allowing quick access to a contiguous block of rows.

**Example**:
```sql
SELECT * FROM Employees WHERE HireDate BETWEEN '2023-01-01' AND '2023-12-31';
```

---

### Summary
- **Full Table Scan**: Scans the entire dataset, no index used, generally slower for large datasets.
- **Point Query**: Uses an index to retrieve specific rows, very fast for indexed columns.
- **Range Query**: Uses a sorted index to retrieve rows within a specified range, efficient for sorted indexes.

These three types of queries form the foundation of SQL operations and cover most of the ways in which data is accessed and manipulated in a database. If you have more questions or need further examples, just let me know! 😊

Got it! Let's explore how hash tables are used for different types of queries. I'll break it down into markdown format for clarity:

---

### 2.2 Using Hash Tables for Different Types of Queries

#### Introduction to Hash Tables
Hash tables (or hash maps) are data structures that allow for fast data retrieval based on a key. They use a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

---

#### 1. Full Table Scan
**Description**: This type of query examines each row in the table, and typically, hash tables are not used because every row needs to be processed.

- **Use Case**: Necessary when the operation requires examining all rows.
- **Example**:
  ```sql
  SELECT * FROM Employees;
  ```
- **Role of Hash Tables**: Not commonly used since the query requires a complete scan of all rows regardless of indexes.

---

#### 2. Point Query
**Description**: This type of query retrieves a specific row or set of rows using a hash table, which is highly efficient for exact key lookups.

- **Use Case**: Efficient for fetching data based on a unique key.
- **Example**:
  ```sql
  SELECT * FROM Employees WHERE EmployeeID = 123;
  ```
- **Role of Hash Tables**:
  - **Index Creation**: Hash tables are used to create an index on the `EmployeeID` column.
  - **Lookup Operation**: The hash function quickly computes the index where the employee with `EmployeeID = 123` is stored, enabling fast retrieval.

---

#### 3. Range Query
**Description**: This type of query retrieves rows within a specified range. Hash tables are not inherently suitable for range queries because they are optimized for exact matches.

- **Use Case**: Useful for retrieving data within a specific range.
- **Example**:
  ```sql
  SELECT * FROM Employees WHERE HireDate BETWEEN '2023-01-01' AND '2023-12-31';
  ```
- **Role of Hash Tables**:
  - **Combination with Other Data Structures**: Range queries are typically better handled by B-trees or similar data structures that maintain sorted order. However, hash tables can be used in conjunction with these structures to enhance performance.
  - **Hybrid Approaches**: Sometimes, hash tables are used to quickly narrow down the potential range of records before a secondary index structure completes the range query.

---

### Summary
- **Full Table Scan**: Hash tables are not commonly used since every row must be scanned.
- **Point Query**: Hash tables provide fast retrieval based on an exact key by using a hash function to compute the index.
- **Range Query**: Hash tables alone are not suitable for range queries, but they can be combined with other data structures to enhance performance.

Hash tables are powerful tools for specific key-based lookups, providing O(1) average time complexity for such operations. They are less suited for operations requiring sorted data or full dataset scans.

If you need more detailed explanations or examples, feel free to ask! 😊
```