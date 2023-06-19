-- Table : Users
CREATE TABLE IF NOT EXISTS Users (
    ID INTEGER PRIMARY KEY,
    Pseudo TEXT NOT NULL,
    Psw TEXT NOT NULL,
    Email TEXT NOT NULL
);

-- Table : Categories
CREATE TABLE IF NOT EXISTS Categories (
    ID INTEGER PRIMARY KEY,
    NameCategories TEXT NOT NULL
);

-- Table : Comments
CREATE TABLE IF NOT EXISTS Comments (
    ID INTEGER PRIMARY KEY,
    ID_Users INTEGER,
    ID_Posts INTEGER,
    MessageContent TEXT,
    DateComment DATETIME,
    FOREIGN KEY (ID_Users) REFERENCES Users(ID),
    FOREIGN KEY (ID_Posts) REFERENCES Posts(ID)
);

-- Table : Likes
CREATE TABLE IF NOT EXISTS Likes (
    ID INTEGER PRIMARY KEY,
    ID_Comments INTEGER,
    ID_Users INTEGER,
    ID_Posts INTEGER,
    DateLike DATETIME,
    FOREIGN KEY (ID_Comments) REFERENCES Comments(ID),
    FOREIGN KEY (ID_Users) REFERENCES Users(ID),
    FOREIGN KEY (ID_Posts) REFERENCES Posts(ID)
);

-- Table : Posts
CREATE TABLE IF NOT EXISTS Posts (
    ID INTEGER PRIMARY KEY,
    ID_Users INTEGER,
    ID_Comments INTEGER,
    ID_Categories INTEGER,
    ID_Likes INTEGER,
    FOREIGN KEY (ID_Users) REFERENCES Users(ID),
    FOREIGN KEY (ID_Comments) REFERENCES Comments(ID),
    FOREIGN KEY (ID_Categories) REFERENCES Categories(ID),
    FOREIGN KEY (ID_Likes) REFERENCES Likes(ID)
);
