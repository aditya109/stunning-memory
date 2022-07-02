FROM ruby
WORKDIR /app
COPY Gemfile Gemfile.lock ./
RUN bundle install
COPY app.rb config.ru ./
EXPOSE 3000
CMD ["bundle", "exec", "rackup", "--host", "0.0.0.0", "-p", "3000"]
