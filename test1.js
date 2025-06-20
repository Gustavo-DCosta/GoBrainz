import { Client } from "@notionhq/client"
import { config } from "dotenv"

config()

const pageId = process.env.DataBase_URL
const apiKey = process.env.NotionWorspace

const notion = new Client({ auth: apiKey })

/* 
---------------------------------------------------------------------------
*/

/**
 * Resources:
 * - Appending block children endpoint (notion.blocks.children.append(): https://developers.notion.com/reference/patch-block-children)
 * - Working with page content guide: https://developers.notion.com/docs/working-with-page-content
 */

async function main() {
  const blockId = pageId // Blocks can be appended to other blocks *or* pages. Therefore, a page ID can be used for the block_id parameter
  const newHeadingResponse = await notion.blocks.children.append({
    block_id: blockId,
    // Pass an array of blocks to append to the page: https://developers.notion.com/reference/block#block-type-objects
    children: [
      {
        heading_2: {
          rich_text: [
            {
              text: {
                content: "This worked", // This is the text that will be displayed in Notion
              },
            },
          ],
        },
      },
    ],
  })

  // Print the new block(s) response
  console.log(newHeadingResponse)
}

main()