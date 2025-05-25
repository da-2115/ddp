//
//  ViewScores.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 24/5/2025.
//

import SwiftUI

struct ViewScores: View {
    @StateObject private var viewModel = ViewModel()
    @State private var showEnd = false
    @State private var endToDisplay = 1
    @State private var visible = true
    
    var body: some View {
        
        if (visible)
        {
            VStack {
                        Text("Scores").font(.largeTitle).bold()

                        // Table-like layout using List
                        List {
                            HStack {
                                Text("ScoreID").bold()
                                Spacer()
                                Text("Score").bold()
                                Spacer()
                                Text("EndID").bold()
                                Spacer()
                                Text("ArrowNumber").bold()
                            }

                            ForEach(viewModel.scores) { score in
                                HStack {
                                    Text("\(Int(score.ScoreID ?? 0))")
                                    Spacer()
                                    Text("\(score.Score)")
                                    Spacer()
                                    Text("\(score.EndID)")
                                    Spacer()
                                    Text("\(score.ArrowNumber)")
                                    
                                    Button("View End", action: {
                                        viewEnd(endId: score.EndID)
                                    })
                                }
                            }
                        }
                
                Button("Add Scores", action: {
                    viewAddScores()
                }).buttonStyle(.borderedProminent)
                    .tint(.green)
                    .controlSize(.large)
                    }
                    .onAppear {
                        viewModel.fetchScores()
                    }
        }
        else if (showEnd)
        {
            ViewEnd(endToDisplay: $endToDisplay)
        }
        else
        {
            AddScores()
        }
        
    }
    
    private func viewEnd(endId: Int) -> Void
    {
        endToDisplay = endId
        showEnd = true
        visible = false
    }
    
    private func viewAddScores() -> Void
    {
        visible = false
    }
}
