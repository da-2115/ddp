//
//  ContentView.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 20/5/2025.
//

import SwiftUI

struct ContentView: View
{
    @StateObject private var viewModel = ViewModel()

    var body: some View
    {
        VStack
        {
            Login()
        }
    }
}
