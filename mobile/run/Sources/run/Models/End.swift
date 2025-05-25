/*
    EndID | RangeID | ArcheryAustraliaID | FinalScore | Staged
*/

import Vapor

struct End : Content
{
    let EndID: Int
    let RangeID: Int
    let ArcheryAustraliaID: String
    let FinalScore: Int
    let Staged: Bool
}