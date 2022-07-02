require 'sinatra'
require 'socket'

set :port, 3000

get '/welcome' do
  puts '/welcome was hit successfully !'
  "Welcome to server ! ⚡ ⛈️"
end

get '/status' do
  puts '/status was hit successfully !'
  "The server is running ! 🏃‍♂️⏩"
end

get '/ping' do
  hostname = Socket.gethostname
  puts '/ping was hit successfully !'
  "Hi there ! 👋 My host-id is #{hostname} 🕸️"
end

