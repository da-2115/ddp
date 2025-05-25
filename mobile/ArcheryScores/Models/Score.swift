//
//  Score.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 24/5/2025.
//

import Vapor

struct Score: Content, Codable, Identifiable
{
    var id: Int? { ScoreID }
    
    let ScoreID: Int?
    let EndID: Int
    let ArrowNumber: Int
    let Score: String
}
