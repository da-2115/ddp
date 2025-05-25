//
//  Member.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 24/5/2025.
//

import Vapor

struct Member: Content, Codable, Identifiable {
    var id: String { ArcheryAustraliaID }

    let ArcheryAustraliaID: String
    let FirstName: String
    let Gender: String
    let PasswordHash: String
    let DateOfBirth: Date
    let ClubRecorder: Bool
}
