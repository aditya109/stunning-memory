package com.github.app.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.net.InetAddress;
import java.net.UnknownHostException;

@RestController
public class RouteController {
    @GetMapping("/")
    public String home() {
        System.out.println("/home was hit successfully !");
        return "Server is running";
    }

    @GetMapping("/welcome")
    public String welcome() {
        System.out.println("/welcome was hit successfully !");
        return "Welcome to server ! ⚡ ⛈";
    }

    @GetMapping("/status")
    public String status() {
        System.out.println("/status was hit successfully !");
        return "The server is running ! \uD83C\uDFC3\u200D♂️⏩";
    }

    @GetMapping("/ping")
    public String ping() {
        System.out.println("/ping was hit successfully !");
        InetAddress ip;
        String hostname = "";
        try {
            ip = InetAddress.getLocalHost();
            hostname = ip.getHostName();
            return "Hi there ! 👋 My host-id is " + hostname + " 🕸";
        } catch (UnknownHostException e) {
            e.printStackTrace();
        }

        return hostname;
    }

}