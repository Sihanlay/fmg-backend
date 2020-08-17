package mailUtils

import (
	"strings"
)

const template = `
<!DOCTYPE html>
<html lang="en" style="margin: 0; background: #f5f5f5 !important">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>邮箱验证</title>
</head>
<body>
    <table style=" border-collapse: collapse; 
        width: 540px; 
        margin: 0 auto; 
        top: 100px; 
        margin-bottom: 100px;
        position: relative; 
        background: white; 
        border-radius: 4px; 
        box-shadow: 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 9px 28px 8px rgba(0, 0, 0, 0.05)
    ">
        <tbody>
            <tr>
                <td style="font-size: 32px; 
                    line-height: 36px; 
                    height: 40px; 
                    background: #40a9ff; 
                    color: white; 
                    border-radius: 4px 4px 0 0; 
                    padding: 15px 24px 8px;
                    margin-block-start: 0.67em;
                    margin-block-end: 0.67em;
                    margin-inline-start: 0px;
                    margin-inline-end: 0px;
                    font-weight: bold;
                ">
				</td>
            </tr>
            <tr>
                <td style="padding-top: 24px; 
                    text-align: center; 
                    line-height: 32px; 
                    font-size: 16px;
                    margin-block-start: 1em;
                    margin-block-end: 1em;
                    margin-inline-start: 0px;
                    margin-inline-end: 0px;
                ">
                    欢迎使用wzlz-backend，请验证邮箱。
					<br/>
					如果此操作非您本人发起的，请忽略此邮件。
                </td>
            </tr>
            <tr>
                <td style="text-align: center; 
                    font-size: 28px;
                    line-height: 30px; 
                    padding: 32px 0;
                    letter-spacing: 4px;
                    margin-block-start: 0.83em;
                    margin-block-end: 0.83em;
                    margin-inline-start: 0px;
                    margin-inline-end: 0px;
                    font-weight: bold;
                ">
                    /*token*/
                </td>
            </tr>
            <tr>
                <td style="
                    text-align: center;
                    font-size: 14px;
                    color: rgb(245, 34, 45);
                    margin: 0px 24px;
                    padding: 24px 0px 32px;
                    line-height: 28px;
                    margin-block-start: 1em;
                    margin-block-end: 1em;
                    margin-inline-start: 0px;
                    margin-inline-end: 0px;
                ">
                    验证码2小时内有效，请尽快完成验证操作。
                </td>
            </tr>
            <tr>
                <td style="
                    margin: 0px 24px;
                    border-top: 1px solid rgb(240, 240, 240);
                    padding: 12px 0px 16px;
                    text-align: center;
                    color: rgb(191, 191, 191);
                    font-size: 14px;
                ">
                    本邮件由系统自动发出，请勿回复！
                </td>
            </tr>
        </tbody>
    </table>
</body>
</html>
`


func generateBody(token string) string {
	return strings.Replace(template, "/*token*/", token, -1)
}
