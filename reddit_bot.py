#!/usr/bin/env python
# -*- coding: utf8 -*- 

import praw
import config

shrug = '¯\_(ツ)_/¯'

def bot_login():
    r = praw.Reddit(username = config.username,
			password = config.password,
			client_id = config.client_id,
			client_secret = config.client_secret,
			user_agent = "Jizzy_Gillespie92 limb retrieval bot")
    return r

def run_bot(r):
	print shrug
	for comment in r.subreddit('test').comments(limit=105):
		if shrug.decode('utf-8') in comment.body:
			print "shrug found!", shrug
			comment.reply("You dropped this \ ")
r = bot_login()
run_bot(r)