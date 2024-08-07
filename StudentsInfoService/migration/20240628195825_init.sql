-- +goose Up
-- +goose StatementBegin
CREATE TABLE Faculty (
    facultyId   SERIAL PRIMARY KEY,
    facultyName TEXT   NOT NULL CHECK(facultyName != '')
);

CREATE TABLE Hostel (
    hostelId    SERIAL PRIMARY KEY,
    hostelName  TEXT   NOT NULL CHECK(hostelName != '')
);

CREATE TABLE Student (
    id          SERIAL  PRIMARY KEY,
    secondName  TEXT    NOT NULL CHECK(secondName != ''),
    firstName   TEXT    NOT NULL CHECK(firstName != ''),
    age         INTEGER NOT NULL CHECK(age >= 1),
    faculty_id  INTEGER REFERENCES Faculty(facultyId),
    hostel_id   INTEGER REFERENCES Hostel(hostelId),
    room        INTEGER NOT NULL CHECK(room >= 1)
);

INSERT INTO Hostel (hostelName)
VALUES
    ('1'),
    ('2'),
    ('3'),
    ('4'),
    ('5'),
    ('6'),
    ('7'),
    ('8'),
    ('9'),
    ('10'),
    ('11'),
    ('12');

INSERT INTO Faculty (facultyName)
VALUES
    ('ФИВТ'),
    ('ФРТК'),
    ('ФАКИ'),
    ('ФУПМ'),
    ('ФБМФ'),
    ('ФАЛТ'),
    ('ФФКЭ');

INSERT INTO Student (secondName, firstName, age, faculty_id, hostel_id, room)
VALUES ('Вербов', 'Ярослав', 22, 1, 9, 41);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Student;
DROP TABLE Faculty;
DROP TABLE Hostel;
-- +goose StatementEnd
