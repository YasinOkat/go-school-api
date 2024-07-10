CREATE DATABASE school;
USE school;

CREATE TABLE user_type
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO user_type (name)
VALUES ('student'),
       ('lecturer'),
       ('admin');

CREATE TABLE major
(
    id            INT PRIMARY KEY AUTO_INCREMENT,
    name          VARCHAR(255) NOT NULL,
    department_id INT          NOT NULL
);


CREATE TABLE user
(
    id           INT PRIMARY KEY AUTO_INCREMENT,
    username     VARCHAR(255) UNIQUE NOT NULL,
    password     VARCHAR(255)        NOT NULL,
    first_name   VARCHAR(50)         NOT NULL,
    last_name    VARCHAR(50)         NOT NULL,
    phone_number VARCHAR(50)         NOT NULL,
    email        VARCHAR(255)        NOT NULL,
    user_type_id INT                 NOT NULL,
    status       BOOL                NOT NULL DEFAULT 1,
    FOREIGN KEY (user_type_id) REFERENCES user_type (id),
    INDEX (user_type_id),
    INDEX (username)
);

CREATE TABLE student
(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    user_id  INT UNIQUE NOT NULL,
    major_id INT        NOT NULL,
    status   BOOL       NOT NULL DEFAULT 1,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (major_id) REFERENCES major (id),
    INDEX (user_id),
    INDEX (major_id)
);

CREATE TABLE course
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

CREATE TABLE major_course
(
    major_id  INT NOT NULL,
    course_id INT NOT NULL,
    FOREIGN KEY (major_id) REFERENCES major (id),
    FOREIGN KEY (course_id) REFERENCES course (id),
    PRIMARY KEY (major_id, course_id)
);

CREATE TABLE student_course
(
    student_id INT NOT NULL,
    course_id  INT NOT NULL,
    FOREIGN KEY (student_id) REFERENCES student (id),
    FOREIGN KEY (course_id) REFERENCES course (id),
    PRIMARY KEY (student_id, course_id)
);

CREATE TABLE student_grade
(
    student_id INT           NOT NULL,
    course_id  INT           NOT NULL,
    grade      DECIMAL(4, 2) NULL,
    FOREIGN KEY (student_id) REFERENCES student (id),
    FOREIGN KEY (course_id) REFERENCES course (id),
    PRIMARY KEY (student_id, course_id)
)