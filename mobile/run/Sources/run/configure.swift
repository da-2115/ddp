@preconcurrency import MySQLKit
@preconcurrency import AsyncKit
import Vapor

struct MySQLStorageKey: StorageKey {
    typealias Value = EventLoopGroupConnectionPool<MySQLConnectionSource>
}

public func configure(_ app: Application) throws {
    // Setup MySQL config
    let mysqlConfig = MySQLConfiguration(
        hostname: "127.0.0.1", // Use "mariadb" if needed in Docker
        port: 3306,
        username: "root",
        password: "1234",
        database: "ARCHERYDB",
        tlsConfiguration: .forClient(certificateVerification: .none) // ✅ Disabled verification for dev
    )

    let source = MySQLConnectionSource(configuration: mysqlConfig)
    let pool = EventLoopGroupConnectionPool(source: source, on: app.eventLoopGroup)

    app.storage[MySQLStorageKey.self] = pool

    // ✅ Properly shut down pool when app stops
    app.lifecycle.use(MySQLPoolLifecycleHandler(pool: pool))

    try routes(app)
}

private struct MySQLPoolLifecycleHandler: LifecycleHandler {
    let pool: EventLoopGroupConnectionPool<MySQLConnectionSource>

    func shutdown(_ application: Application) {
        pool.shutdown()
    }
}
