This is a webservice using following
 - golang's net/http capabilities
 - golang's html/template capabilities
 - go-sqlite driver, which connects to sqlite
 - gorrila mux

This service is used by an ios client making REST calls
to this. Although, can also be invoked from any client making
REST calls.

Supported REST APIs
The root is http://localhost:8080/
Supported paths are
 "/prospect/view/"
 "/prospect/view/{criteria}"
 "/prospect/add/"
 "/prospect/update/"

 "/participant/add/"
 "/participant/view/"
 "/participant/view/userid/{userid}"
 "/participant/view/prospectid/{id:[0-9]+}"
 "/participant/update/"

 "/discussion/add/"
 "/discussion/view/"
 "/discussion/view/prospectid/{id:[0-9]+}"
 "/discussion/update/"
 "/discussion/view/html/"

 Table Schema

 CREATE TABLE "prospects" (
  `ProspectID`INTEGER PRIMARY KEY AUTOINCREMENT,
  `Name`TEXT,
  `ConfDate`TEXT,
  `TechStack`TEXT,
  `Domain`TEXT,
  `DesiredTeamSize`INT,
  `Notes`TEXT,
  `SalesID`INT,
  `CreateDate`TEXT,
  `StartDate`TEXT,
  `BUHead`TEXT,
 `TeamSize`INT);

CREATE TABLE "participants" (
`ProspectID` INTEGER,
`UserID` TEXT,
`Included` TEXT,
`Participation` TEXT,
 FOREIGN KEY(ProspectID) REFERENCES prospects(ProspectID));

CREATE TABLE "discussions" (
`DiscussionID` INTEGER PRIMARY KEY AUTOINCREMENT,
`ProspectID` INTEGER,
`UserID` TEXT,
`Query` TEXT,
`Answer` TEXT,
 FOREIGN KEY(ProspectID) REFERENCES prospects(ProspectID));