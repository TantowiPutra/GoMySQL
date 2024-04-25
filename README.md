# GoMySQL
Repository to Learn MySQL database connectivity using Golang
<!-- Database packages -->
- By default has a package called database, Filled with standard interface to interact with the database
- Which let us to access any type of database using the same code
- The sql code might difer, depending on the type of database we're using

How the package works

application -> database interface -> database driver -> DBMS

To interact with the database, we need to install database driver. Since database package only contain interfaces, the program wouldn't understand the implementation without the corresponding package.
For example, we can install database driver from: https://golang.org/s/sqldrivers

<!-- Database Pooling -->
- sql.DB isn't a connection to database
- Instead, it's the creation of pool to the database, known as "Database Pooling"
- In sql.DB, Golang manages database connection automatically, therefre, we don't need to manage database connection manually
- Database Pooling has several advantages, such as: determine minium and maximum database connection that could be made by Golang, so that it's not flooding the database connection, because, most of the time, there's a limit on the amount of connections that can be made by the corresponding database 

<!-- Several Settings for Database Pooling -->
Method 
1. (DB) SetMaxIdleConns(number) -> set the minimum amount of connection to be made
2. (DB) SetMaxOpenConns(number) -> set the maximum amount of connection to be made
3. (DB) SetConnMaxIdleTime(duration) -> set the duration before an idle connection is terminated
3. (DB) SetConnMaxLifetime(duration) -> set the duration how long a connection can be used

<!-- Execute SQL Query -->
1. (DB) ExecContext(context, sql, param) -> RUn Query that doesn't return rows

<!-- What is Database Pooling ? -->
1. Database Pooling is pre-established database connection, in which the database connection is already declared before receiving request from a user
2. Using database pooling, we can define the minimum and maximum limit of database connection
3. Normally, without using database pooling, each request will create a database connection. This approach is inneficient as it will exhaust the system memory exponentially with the number of requests