import logging
import requests

"""
checkFrameworkVersion checks the python-telegram-bot framework version and ensures the version used is updated enough.
"""
def checkFrameworkVersion():
    try:
        from telegram import __version_info__
    except ImportError:
        __version_info__ = (0, 0, 0, 0, 0)

    if __version_info__ < (20, 0, 0, "alpha", 1):
        raise RuntimeError("python-telegram-bot version is outdated")
    logging.info(f"[framework version check]: success\n\
        python-telegram-bot version: {__version_info__}")


"""
executeHealthCheck makes a HTTPS getMe call using the Telegram Bot API, to validate the functionality of the provided token.
"""
def executeHealthCheck(bot_token):
    response = requests.get(f"https://api.telegram.org/bot{bot_token}/getMe")
    if response.status_code == 401:
        logging.fatal(
            f"[bb8 health check] error\n\
            code: {response.status_code}\n\
            reason: invalid_bot_token\
            content: {response.content}")
        raise RuntimeError("invalid bot token")
    if response.status_code != 200:
        logging.fatal(
            f"[bb8 health check] error\n\
            code: {response.status_code}\n\
            content: {response.content}")
        raise RuntimeError("unexpected health check failure")
    else:
        logging.info(
            f"[bb8 health check] success\n\
            code: {response.status_code}\n\
            content: {response.content}")
