#!/usr/bin/env python
# -*- coding: utf8 -*- 

import praw
import config
import footer

shrug = '¯\_(ツ)_/¯'
shoulders = '¯\\\_(ツ)_/¯'
decapitated = '¯\ _(ツ)_/¯'

def bot_login():
    r = praw.Reddit(username = config.username,
			password = config.password,
			client_id = config.client_id,
			client_secret = config.client_secret,
			user_agent = "LimbRetrieval-Bot by u/Jizzy_Gillespie92")
    return r

def run_bot(r):
	subreddit = r.subreddit('all')
	try:
		comments = subreddit.stream.comments()
		for comment in comments:
			text = comment.body
			author = comment.author

			if shrug.decode('utf-8') in text: #decode required for unicode characters
			#create/open log.txt
				try:
					file = open("log.txt", "a")
				except IOError:
					file = open("log.txt", "w")

				file.write(str(author) + '\n')
				file.close()

				# print "Missing limb found in: ", "(",author,")", text
				try:
					comment.reply("You dropped this \ " + footer.footerComment)
				except:
					file = open("log.txt", "w")
					file.write("Exception caught!" + '\n') #TODO error details
					file.close()
			
			elif shoulders.decode('utf-8') in text:
				try:
					file = open("log.txt", "a")
				except IOError:
					file = open("log.txt", "w")

				file.write(str(author) + '\n')
				file.close()
				# print "Shoulders lost forever: ", "(",author,")", text
				try:
					comment.reply("I have retrieved these for you _ _" + footer.footerComment)
				except:
					file = open("log.txt", "w")
					file.write("Exception caught!" + '\n') #TODO error details
					file.close()					
	except:
		try:
			file = open("log.txt", "a")
		except IOError:
			file = open("log.txt", "w")
			file.write("Exception caught!" + '\n') #TODO error details
		file.close()

r = bot_login()
run_bot(r)