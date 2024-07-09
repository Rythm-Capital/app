import styled, { css } from 'styled-components'


const Box = styled.div`
    border-radius: 2px;
    z-index: ${props => {
        let zoffset = (props.level / 100)
        let z = 11 - zoffset
        return z
    }};
    height: ${props => {
        let sizeOffset = 1000 - props.level 
        let pixelOffset = 80 * (sizeOffset / 100)
        let height = 800 - pixelOffset
        let size = height/3 
        return `${size}px;`
    }}
    width: ${props => {
        let sizeOffset = 1000 - props.level 
        let pixelOffset = 80 * (sizeOffset / 100)
        let height = 800 - pixelOffset
        let size = height/3 
        return `${size}px;`
    }};
    // position: absolute;

    grid-row: ${props => {
        
        let index = 10 - props.level / 100
        
        let item = props.glyph[index]
        

        if (props.up) {

            if (props.level === 1000) {
                return `1`
            }

            const lastBoxIndex = props.glyph.findLastIndex(el =>  el.level > item.level) 
            const lastBox = props.glyph[lastBoxIndex]
            if(!lastBox) {
                return `${item.row -1}`
            }
            
            props.glyph[index].row = lastBox.row
            
            return  `${lastBox.row}`  

        } if (!props.up) {
            
            if (props.level === 1000) {
                return `1`
            }

            const lastBoxIndex = props.glyph.findLastIndex(el =>  el.level > item.level ) 
            const lastBox = props.glyph[lastBoxIndex]
            
            if(!lastBox) {
                return `${item.row}`
            }
            
            props.glyph[index].row  = lastBox.row + 1
            
            return  `${lastBox.row + 1}`
        }
        
    }};
    grid-column-start: ${props => {
        if (props.level == 1000) {
            return 1
        } else {
            let offset = 1000 - props.level
            let row = (offset / 100) + 1
            return `${row}`
        }
    }};
    
    background: ${props => props.up ? "#104B0051" : "#55000A51"};
    border:1px solid ${props => props.up ? "darkgreen" : "maroon"};
    `
    
    // bottom: calc(80px/3);

    
export default Box