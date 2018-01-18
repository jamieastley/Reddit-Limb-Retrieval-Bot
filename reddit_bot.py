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
			user_agent = "LimbRetrieval-Bot by u/Jizzy_Gillespie92")
    return r

def run_bot(r):
	subreddit = r.subreddit('all')
	comments = subreddit.stream.comments()
	for comment in comments:
		text = comment.body
		author = comment.author
		if shrug.decode('utf-8') in text:
			print "shrug found in- ", author, text
			print(r.user.me())
			comment.reply("You dropped this \ ")
			
			# TODO write to log all occurrences
r = bot_login()
run_bot(r)