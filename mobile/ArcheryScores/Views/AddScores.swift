//
//  AddScores.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 24/5/2025.
//

import SwiftUI

struct AddScores: View {
    @StateObject private var viewModel = ViewModel()
    
    // Use String bindings for text fields
    @State private var endIDText: String = "1"
    @State private var arrowNumberText: String = "1"
    @State private var scoreText: String = ""
    @State private var visible = true

    var body: some View {
        if ( visible )
        {
            Form {
                
                    TextField("End ID", text: $endIDText)
                        .keyboardType(.numberPad)

                    TextField("Arrow Number", text: $arrowNumberText)
                        .keyboardType(.numberPad)

                    TextField("Score (e.g., X, 10, 9...)", text: $scoreText)

                    Button("Submit") {
      
                        let newScore = Score(
                            ScoreID: nil,
                            EndID: Int(endIDText) ?? 0,
                            ArrowNumber: Int(arrowNumberText) ?? 0,
                            Score: scoreText
                        )
                        addScore(newScore: newScore)
                    }
                }

                if let error = viewModel.errorMessage {
                    Text("Error: \(error)")
                        .foregroundColor(.red)
                
            }
        }
        
        else
        {
            ViewScores()
        }
        
    }
    
    private func addScore(newScore: Score) -> Void
    {
        viewModel.postScore(newScore)
        visible = false
    }
}
