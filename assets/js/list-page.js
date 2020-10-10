/** 
 * list-page.js 
 * this script from http://book.jirengu.com/jirengu-inc/jrg-renwu10/homework/%E5%AD%99%E7%BA%A2%E7%85%A7/mission30/demo.html
 *
 */


var waterBasic = (function(){
    function init(){
        /* get container width */
        var minWidth = 504;
        var containerWidth = $(window).width();
        for(;;){
            if(containerWidth > minWidth+252){
                minWidth += 252;
            } else {
                break;
            }
        }
        console.log(minWidth)
        /* get node width */
        $(".water-basic").width(minWidth);
        var nodeWidth = $(".item").outerWidth(true),
            colNum = parseInt( $(".water-basic").width() / nodeWidth ),
            colSumHeight = [];
        for (var i=0;i<colNum;i++) {
            colSumHeight.push(0);
        }
        $(".item").each(function(){
            var $cur = $(this),
                idx = 0,
                minSumHeight = colSumHeight[0];
            // 获取到solSumHeight中的最小高度
            for (var i=0;i<colSumHeight.length;i++) {
                if (minSumHeight > colSumHeight[i]) {
                    minSumHeight = colSumHeight[i];
                    idx = i;
                }
            }
            // 设置各个item的css属性
            $cur.css({
                left: nodeWidth*idx,
                top: minSumHeight
            })
            // 更新solSumHeight
            colSumHeight[idx] = colSumHeight[idx] + $cur.outerHeight(true);
        })
    }
    // 设置窗口改变时也能重新加载
    $(window).on("resize", function(){
        init();
    })
    return {
        init: init
    }
})();

waterBasic.init();

