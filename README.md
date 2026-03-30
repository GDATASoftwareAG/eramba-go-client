# eramba go client

> [!IMPORTANT]
> This is a early version not all API's are supported and tested.

## Supported Eramba Instance

Minimal requirement is Eramba 3.28.0.

## Supported APIs

| Type                      | Methods                         | Custom Fields | Comments & Attachments |
|---------------------------|---------------------------------|---------------|------------------------|
| Assets                    | Get, Index, Patch, Post, Delete |               |                        |
| Asset Reviews             | Index                           |               |                        |
| Groups                    | Get, Index                      |               |                        |
| Users                     | Get, Index                      |               |                        |
| Security Policies         | Get, Index, Patch, Post         |               |                        |
| Security Policy Reviews   | Get, Index, Patch, Post, Delete |               |                        |
| Projects                  | Index, Patch, Post              | Supported     | Comments               |
| Project Comments          | Post                            |               |                        |
| Risk Exceptions           | Get, Index, Patch, Post         |               |                        |
| Risks                     | Get, Index, Patch, Post         | Supported     | Comments               |
| Risk Reviews              | Index, Patch, Post              |               |                        |
| Risk Threats              | Get, Index                      |               |                        |
| Security Services         | Get, Index, Patch               |               | Comments               |
| Security Service Comments | Index                           |               |                        |
| Third Parties             | Get, Index, Patch               | Supported     |                        |
| Third Party Risks         | Get, Index, Patch, Post         | Supported     |                        |
| Third Party Risk Reviews  | Index, Post                     |               |                        |

## Disclaimer

**This project is not associated with Eramba LTD.**