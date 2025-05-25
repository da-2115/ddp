import Vapor

struct Score: Content 
{
    let ScoreID: Int?
    let EndID: Int
    let ArrowNumber: Int
    let Score: String
}