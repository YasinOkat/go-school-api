# Go School Management API

This is a REST API for a school management system built using Go, Gin framework, and JWT for authentication. The API allows for user management, course management, and student-course selections with appropriate access controls based on user roles.

The API connects to MySQL database, the SQL code to create the tables can be accessed at db.sql file. No ORM is used, all queries are raw SQL queries.

## Features

- **User Management**: Create and manage users (students, lecturers, admins).
- **Course Management**: Manage courses available in the school.
- **Student-Course Selection**: Students can select courses.
- **Authentication**: JWT-based authentication.
- **Role-Based Access Control**: Different access levels for students, lecturers, and admins.

## Documentation

- **Swagger**: The swagger documentation can be initialized git `swag init` command. Go the the /swagger/index.html.
