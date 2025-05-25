//
//  ViewEnd.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 25/5/2025.
//

import SwiftUI

struct ViewEnd: View
{
    @StateObject private var viewModel = ViewModel()
    @Binding var endToDisplay: Int
    @State private var end: End?
    @State private var visible: Bool = true
    
    var body: some View
    {
        if ( visible )
        {
            VStack
            {
                Text("End").font(.largeTitle).bold()
                List {
                    Text("End ID: \(end?.EndID ?? 0)").font(.title3).bold()
                    Text("Final Score: \(end?.FinalScore ?? 0)")
                    
                }
                
                    Button("Go Back", action: { goBack() }).buttonStyle(.borderedProminent)
                        .tint(.red)
                        .controlSize(.large)
                
            }
            .onAppear
            {
                viewModel.fetchEnds()
            }
            .onReceive(viewModel.$ends)
            { _ in
                DispatchQueue.main.asyncAfter(deadline: .now() + 0.1)
                {
                    initEnd()
                }
            }
        }
        
        else
        {
            ViewScores()
        }
    }
    
    private func goBack() -> Void
    {
        visible = false
    }
    
    private func initEnd() -> Void {
        if let found = viewModel.ends.first(where: { $0.EndID == endToDisplay })
        {
            end = found
        }
        
        else
        {
            end = nil
        }
    }

}
