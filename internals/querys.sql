CREATE TABLE Users (
                       UserID INT AUTO_INCREMENT PRIMARY KEY,
                       Username VARCHAR(255) NOT NULL,
                       Email VARCHAR(255) NOT NULL,
                       Password VARCHAR(255) NOT NULL,
                       FirstName VARCHAR(50),
                       LastName VARCHAR(50),
                       AvatarURL VARCHAR(255),
                       RegistrationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Projects (
                          ProjectID INT AUTO_INCREMENT PRIMARY KEY,
                          ProjectName VARCHAR(255) NOT NULL,
                          UserID INT,
                          FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Tasks (
                       TaskID INT AUTO_INCREMENT PRIMARY KEY,
                       TaskTitle VARCHAR(255) NOT NULL,
                       TaskDescription TEXT,
                       DueDateTime DATETIME,
                       Priority VARCHAR(20),
                       TaskStatus VARCHAR(20),
                       ProjectID INT,
                       UserID INT,
                       ParentTaskID INT,
                       FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID),
                       FOREIGN KEY (UserID) REFERENCES Users(UserID),
                       FOREIGN KEY (ParentTaskID) REFERENCES Tasks(TaskID)
);

CREATE TABLE Subtasks (
                          SubtaskID INT AUTO_INCREMENT PRIMARY KEY,
                          SubtaskTitle VARCHAR(255) NOT NULL,
                          SubtaskStatus VARCHAR(20),
                          TaskID INT,
                          UserID INT,
                          FOREIGN KEY (TaskID) REFERENCES Tasks(TaskID),
                          FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Labels (
                        LabelID INT AUTO_INCREMENT PRIMARY KEY,
                        LabelName VARCHAR(50) NOT NULL
);

CREATE TABLE TaskLabels (
                            TaskID INT,
                            LabelID INT,
                            PRIMARY KEY (TaskID, LabelID),
                            FOREIGN KEY (TaskID) REFERENCES Tasks(TaskID),
                            FOREIGN KEY (LabelID) REFERENCES Labels(LabelID)
);

CREATE TABLE Comments (
                          CommentID INT AUTO_INCREMENT PRIMARY KEY,
                          CommentText TEXT NOT NULL,
                          TaskID INT,
                          UserID INT,
                          Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (TaskID) REFERENCES Tasks(TaskID),
                          FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Reminders (
                           ReminderID INT AUTO_INCREMENT PRIMARY KEY,
                           ReminderDateTime DATETIME NOT NULL,
                           TaskID INT,
                           UserID INT,
                           FOREIGN KEY (TaskID) REFERENCES Tasks(TaskID),
                           FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Collaborators (
                               CollaborationID INT AUTO_INCREMENT PRIMARY KEY,
                               TaskID INT,
                               UserID INT,
                               FOREIGN KEY (TaskID) REFERENCES Tasks(TaskID),
                               FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Notifications (
                               NotificationID INT AUTO_INCREMENT PRIMARY KEY,
                               UserID INT,
                               NotificationType VARCHAR(50) NOT NULL,
                               NotificationContent TEXT,
                               Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE UserPreferences (
                                 UserID INT,
                                 ThemePreferences VARCHAR(50),
                                 NotificationSettings TEXT,
                                 DefaultProjectSettings INT,
                                 OtherPreferences TEXT,
                                 FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE ArchivedTasks (
                               TaskID INT,
                               UserID INT,
                               ArchiveDate DATETIME,
                               IsArchived BOOLEAN,
                               FOREIGN KEY (TaskID) REFERENCES Tasks(TaskID),
                               FOREIGN KEY (UserID) REFERENCES Users(UserID)
);
