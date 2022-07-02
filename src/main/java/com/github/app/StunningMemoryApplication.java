package com.github.app;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;

@SpringBootApplication(exclude = SecurityAutoConfiguration.class)
public class StunningMemoryApplication {

	public static void main(String[] args) {
		SpringApplication.run(StunningMemoryApplication.class, args);
	}

}

