//
//  ViewModel.swift
//  ArcheryScores
//
//  Created by Dylan Armstrong on 24/5/2025.
//

import Foundation
import Combine

class ViewModel: ObservableObject {
    @Published var members: [Member] = []
    @Published var scores: [Score] = []
    @Published var ends: [End] = []
    
    private var cancellables = Set<AnyCancellable>()
    
    @Published var errorMessage: String?

    func fetchMembers() {
        guard let url = URL(string: "http://127.0.0.1:8080/members") else {
            errorMessage = "Invalid URL"
            print("Invalid URL")
            return
        }
        
        print("Fetching members from \(url)")
        
        let task = URLSession.shared.dataTask(with: url) { data, response, error in
            DispatchQueue.main.async {
                if let error = error {
                    self.errorMessage = "Network error: \(error.localizedDescription)"
                    print("Network error: \(error.localizedDescription)")
                    return
                }
                
                guard let data = data else {
                    self.errorMessage = "No data returned"
                    print("No data returned")
                    return
                }
                
                // Optionally print raw JSON to debug
                if let jsonString = String(data: data, encoding: .utf8) {
                    print("Received JSON: \(jsonString)")
                }
                
                do {
                    let decoder = JSONDecoder()
                    decoder.dateDecodingStrategy = .iso8601
                    self.members = try decoder.decode([Member].self, from: data)
                    print("Decoded \(self.members.count) members")
                } catch {
                    self.errorMessage = "Decoding error: \(error.localizedDescription)"
                    print("Decoding error: \(error.localizedDescription)")
                }
            }
        }
        
        task.resume()
    }
    
    func fetchScores() {
        guard let url = URL(string: "http://127.0.0.1:8080/scores") else {
            errorMessage = "Invalid URL"
            print("Invalid URL")
            return
        }
        
        print("Fetching members from \(url)")
        
        let task = URLSession.shared.dataTask(with: url) { data, response, error in
            DispatchQueue.main.async {
                if let error = error {
                    self.errorMessage = "Network error: \(error.localizedDescription)"
                    print("Network error: \(error.localizedDescription)")
                    return
                }
                
                guard let data = data else {
                    self.errorMessage = "No data returned"
                    print("No data returned")
                    return
                }
                
                // Optionally print raw JSON to debug
                if let jsonString = String(data: data, encoding: .utf8) {
                    print("Received JSON: \(jsonString)")
                }
                
                do {
                    let decoder = JSONDecoder()
                    decoder.dateDecodingStrategy = .iso8601
                    self.scores = try decoder.decode([Score].self, from: data)
                    print("Decoded \(self.members.count) members")
                } catch {
                    self.errorMessage = "Decoding error: \(error.localizedDescription)"
                    print("Decoding error: \(error.localizedDescription)")
                }
            }
        }
        
        task.resume()
    }
    
    func fetchEnds() {
        guard let url = URL(string: "http://127.0.0.1:8080/ends") else {
            errorMessage = "Invalid URL"
            print("Invalid URL")
            return
        }
        
        print("Fetching members from \(url)")
        
        let task = URLSession.shared.dataTask(with: url) { data, response, error in
            DispatchQueue.main.async {
                if let error = error {
                    self.errorMessage = "Network error: \(error.localizedDescription)"
                    print("Network error: \(error.localizedDescription)")
                    return
                }
                
                guard let data = data else {
                    self.errorMessage = "No data returned"
                    print("No data returned")
                    return
                }
                
                // Optionally print raw JSON to debug
                if let jsonString = String(data: data, encoding: .utf8) {
                    print("Received JSON: \(jsonString)")
                }
                
                do {
                    let decoder = JSONDecoder()
                    decoder.dateDecodingStrategy = .iso8601
                    self.ends = try decoder.decode([End].self, from: data)
                    print("Decoded \(self.ends.count) ends")
                } catch {
                    self.errorMessage = "Decoding error: \(error.localizedDescription)"
                    print("Decoding error: \(error.localizedDescription)")
                }
            }
        }
        
        task.resume()
    }
    
    func postScore(_ score: Score) {
            guard let url = URL(string: "http://127.0.0.1:8080/scores") else {
                self.errorMessage = "Invalid URL"
                return
            }

            // Prepare the request
            var request = URLRequest(url: url)
            request.httpMethod = "POST"
            request.setValue("application/json", forHTTPHeaderField: "Content-Type")

            // Encode the score to JSON
            do {
                let jsonData = try JSONEncoder().encode(score)
                request.httpBody = jsonData
            } catch {
                self.errorMessage = "Encoding error: \(error.localizedDescription)"
                return
            }

            // Perform the request
            URLSession.shared.dataTaskPublisher(for: request)
                .map(\.data)
                .receive(on: DispatchQueue.main)
                .sink(receiveCompletion: { completion in
                    if case let .failure(error) = completion {
                        self.errorMessage = "Request failed: \(error.localizedDescription)"
                    }
                }, receiveValue: { _ in
                    self.errorMessage = nil // Clear errors on success
                })
                .store(in: &cancellables)
        }

}
