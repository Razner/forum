-- Table : Users
CREATE TABLE Users (
    ID INT PRIMARY KEY,
    Pseudo VARCHAR(24) NOT NULL,
    Psw VARCHAR(50) NOT NULL,
    Email VARCHAR(100) NOT NULL,
);

-- Table : Categories
CREATE TABLE Categories (
    ID INT PRIMARY KEY,
    NameCategories VARCHAR(50) NOT NULL,
);

-- Table : Comments
CREATE TABLE Comments (
    ID INT PRIMARY KEY,
    ID_Users INT,
    ID_Posts INT,
    MessageContent VARCHAR(500),
    DateComment DATETIME,
    FOREIGN KEY (ID_Users) REFERENCES Users(ID),
    FOREIGN KEY (ID_Posts) REFERENCES Posts(ID)
);

-- Table : Likes
CREATE TABLE Likes (
    ID INT PRIMARY KEY,
    ID_Comments INT,
    ID_Users INT,
    ID_Posts INT,
    DateLike TIMESTAMP,
    FOREIGN KEY (ID_Comments) REFERENCES Comments(ID),
    FOREIGN KEY (ID_Users) REFERENCES Users(ID),
    FOREIGN KEY (ID_Posts) REFERENCES Posts(ID)
);

-- Table : Posts 
CREATE TABLE Posts (
    ID INT PRIMARY KEY,
    ID_Users INT,
    ID_Comments INT,
    ID_Categories INT,
    ID_Likes INT,
    FOREIGN KEY (ID_Users) REFERENCES Users(ID),
    FOREIGN KEY (ID_Comments) REFERENCES Comments(ID),
    FOREIGN KEY (ID_Categories) REFERENCES Categories(ID)
    FOREIGN KEY (ID_Likes) REFERENCES Likes(ID)
)