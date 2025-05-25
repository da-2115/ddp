import MySQLKit
import Vapor

func routes(_ app: Application) throws {
    app.get("hello") { req -> String in
        "Hello, world!"
    }

    app.get("members") { req -> EventLoopFuture<[Member]> in
        guard let pool = req.application.storage[MySQLStorageKey.self] else {
            return req.eventLoop.makeFailedFuture(
                Abort(.internalServerError, reason: "MySQL pool not configured")
            )
        }

        return pool.withConnection { conn in
            conn.sql()
                .select()
                .column("*")
                .from("Member")
                .all()
        }.flatMapThrowing { (rows: [any SQLRow]) -> [Member] in
            try rows.map { row in
                return Member(
                    ArcheryAustraliaID: try row.decode(
                        column: "ArcheryAustraliaID", as: String.self),
                    FirstName: try row.decode(column: "FirstName", as: String.self),
                    Gender: try row.decode(column: "Gender", as: String.self),
                    PasswordHash: try row.decode(column: "PasswordHash", as: String.self),
                    DateOfBirth: try row.decode(column: "DateOfBirth", as: Date.self),
                    ClubRecorder: try row.decode(column: "ClubRecorder", as: Bool.self)
                )
            }
        }
    }

    app.get("scores") { req -> EventLoopFuture<[Score]> in
        guard let pool = req.application.storage[MySQLStorageKey.self] else {
            return req.eventLoop.makeFailedFuture(
                Abort(.internalServerError, reason: "MySQL pool not configured")
            )
        }

        return pool.withConnection { conn in
            conn.sql()
                .select()
                .column("*")
                .from("Score")
                .all()
        }.flatMapThrowing { rows in
            try rows.map { row in
                return Score(
                    ScoreID: try row.decode(column: "ScoreID", as: Int.self),
                    EndID: try row.decode(column: "EndID", as: Int.self),
                    ArrowNumber: try row.decode(column: "ArrowNumber", as: Int.self),
                    Score: try row.decode(column: "Score", as: String.self)
                )
            }
        }
    }

    app.get("ends") { req -> EventLoopFuture<[End]> in
        guard let pool = req.application.storage[MySQLStorageKey.self] else {
            return req.eventLoop.makeFailedFuture(
                Abort(.internalServerError, reason: "MySQL pool not configured")
            )
        }

        return pool.withConnection { conn in
            conn.sql()
                .select()
                .column("*")
                .from("End")
                .all()
        }.flatMapThrowing { rows in
            try rows.map { row in
                return End(
                    EndID: try row.decode(column: "EndID", as: Int.self),
                    RangeID: try row.decode(column: "RangeID", as: Int.self),
                    ArcheryAustraliaID: try row.decode(
                        column: "ArcheryAustraliaID", as: String.self),
                    FinalScore: try row.decode(column: "FinalScore", as: Int.self),
                    Staged: try row.decode(column: "Staged", as: Bool.self)
                )
            }
        }
    }

    app.post("scores") { req -> EventLoopFuture<HTTPStatus> in
        // 1. Ensure DB pool exists
        guard let pool = req.application.storage[MySQLStorageKey.self] else {
            return req.eventLoop.makeFailedFuture(
                Abort(.internalServerError, reason: "MySQL pool not configured")
            )
        }

        // 2. Decode the incoming JSON into a Score object
        let incomingScore = try req.content.decode(NewScore.self)

        // 3. Insert the score into the database
        return pool.withConnection { conn in
            conn.sql()
                .insert(into: "Score")
                .columns("EndID", "ArrowNumber", "Score")
                .values(incomingScore.EndID, incomingScore.ArrowNumber, incomingScore.Score)
                .run()
        }.transform(to: .created)  // Return HTTP 201 Created if successful
    }

}
