import logging
import requests

try:
    from telegram import __version_info__
except ImportError:
    __version_info__ = (0, 0, 0, 0, 0)


""" validateFrameworkVersion validates the python-telegram-bot framework version. """
def validateFrameworkVersion():
    if __version_info__ < (20, 0, 0, "alpha", 1):
        raise RuntimeError(f"python-telegram-bot version is outdated: {__version_info__}")
    logging.info("[framework version validation] success : python-telegram-bot framework version %s",
        __version_info__)


""" executeHealthCheck validates the functionality of the provided bot token. """
def executeHealthCheck(bot_token):
    response = requests.get(f"https://api.telegram.org/bot{bot_token}/getMe")

    if response.status_code != 200:
        logging.fatal("[bb8 health check] error: %d; content: %s", response.status_code, response.content)
        raise RuntimeError(f"unexpected health check failure: {response.status_code}")
    logging.info( "[bb8 health check] 200 OK; content: %s", response.content)
