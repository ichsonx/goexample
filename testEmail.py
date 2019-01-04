
from email.mime.text import MIMEText
from email.header import Header
from smtplib import SMTP_SSL
import argparse

smtp_server = "smtp.qq.com"
port = 465
# sender = "sonxz@qq.com"
sender = "sonxz@qq.com,"
pw = "tmakpyqrpqkkbjfa"
mail_content = "halo test mail..."
mail_subject = "test subject"

def _parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-server",
        help="the smtp server address",
    )
    parser.add_argument(
        "-port",
        help="the smtp server port",
    )
    parser.add_argument(
        "-pw",
        help="the smtp server`s password",
    )
    parser.add_argument(
        "-sender",
        help="the sender`s email",
    )
    parser.add_argument(
        "-subject",
        help="the email subject",
    )
    parser.add_argument(
        "-to",
        help="the email who do you want to send. single email also mutilple email, like 'aaa@aaa.com,bbb@bbb.com'",
    )
    parser.add_argument(
        "-cc",
        help="the email who do you want to CC. single email also mutilple email, like 'aaa@aaa.com,bbb@bbb.com'",
    )
    parser.add_argument(
        "-bcc",
        help="the email who do you want to BCC. single email also mutilple email, like 'aaa@aaa.com,bbb@bbb.com'",
    )
    parser.add_argument(
        "-content",
        help="the email content"
    )
    return parser.parse_args()

if __name__ == '__main__':
    args = _parse_args()

    smtp = SMTP_SSL(args.server, int(args.port))
    smtp.set_debuglevel(1)
    smtp.login(args.sender, args.pw)
    msg = MIMEText(args.content, "plain", "utf-8")
    msg["Subject"] = Header(args.subject, "utf-8")
    msg["From"] = args.sender
    msg["To"] = args.to     #可以是多人，则为'one@qq.com,two@qq.com'
    if args.cc != None:
        msg["Cc"] = args.cc     #同上
    if args.bcc != None:
        msg["BCc"] = args.bcc
    smtp.sendmail(sender, args.to, msg.as_string())
    smtp.quit()

