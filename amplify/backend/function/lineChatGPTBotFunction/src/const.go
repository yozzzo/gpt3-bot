package main

import "os"

const BASE_SECRET_PATH = os.Getenv()
const OPEN_AI_API_KEY = aws_systems_manager.get_secret("{BASE_SECRET_PATH}OPEN_AI_API_KEY")
const LINE_CHANNEL_SECRET = aws_systems_manager.get_secret("{BASE_SECRET_PATH}LINE_CHANNEL_SECRET")
const LINE_CHANNEL_ACCESS_TOKEN = aws_systems_manager.get_secret("{BASE_SECRET_PATH}LINE_CHANNEL_ACCESS_TOKEN")
