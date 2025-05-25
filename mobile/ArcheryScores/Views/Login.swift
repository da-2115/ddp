import SwiftUI
import BCrypt

struct Login: View {
    @StateObject private var viewModel = ViewModel()
    @State private var username: String = ""
    @State private var password: String = ""
    @State private var isAuth: Bool = false
    @State private var errorMessage: String = ""
    @State private var visible: Bool = true
    
    var body: some View
    {
        if ( visible )
        {
            VStack
            {
                TextField("Username", text: $username)
                SecureField("Password", text: $password)
                    .autocapitalization(.none)
                
                Button("Login", action: verifyLogin).buttonStyle(.borderedProminent)
                    .tint(.green)
                    .controlSize(.large)
                
                Text("You are now logged in if this is green")
                    .foregroundColor(isAuth ? .green : .red)
                
                
                
                if ( !errorMessage.isEmpty )
                {
                    Text(errorMessage)
                        .foregroundColor(.orange)
                        .font(.caption)
                }
            }
            .onAppear
            {
                viewModel.fetchMembers()
            }
            .padding()
        }
        else
        {
            if ( isAuth )
            {
                ViewScores()
            }
        }
    }
        

    private func verifyLogin()
    {
        guard let member = viewModel.members.first(where:
        {
            $0.ArcheryAustraliaID.lowercased() == username.lowercased()
        })
        else
        {
            isAuth = false
            errorMessage = "User not found."
            return
        }

        let storedHash = member.PasswordHash.trimmingCharacters(in: .whitespacesAndNewlines)

        do
        {
            let result = try BCrypt.Hash.verify(message: password, matches: storedHash)
            isAuth = result
            visible = false
            errorMessage = result ? "" : "Incorrect password."
        }
        
        catch
        {
            isAuth = false
            errorMessage = "Error during verification: \(error.localizedDescription)"
        }
    }
}
