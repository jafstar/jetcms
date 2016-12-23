//COMMON.js
console.log('common js file');


//READY DOC
$(document).ready(function(){



//LIGHTBOX CLICK
$("#lightbox").on('click',function(event){


      //PREVENT DEFAULT
      event.preventDefault();
      
      $("#lightbox").hide();

});



//FLUSH CACHE
$("#Flush").on('click',function(event){


      //PREVENT DEFAULT
      event.preventDefault();
      
      //GRAB URL
      var url = event.target.href;


        //AJAX
        $.ajax({
           type: "GET",
           url: url,
           success: function(res){
               
              //ALERT
              alertify.success(res.message);
           }

         });


});


/*
   //CLICK DELETE BUTTONS
$("#List").on('click',"tr td a.deleteItem",function(event){
      



      //GRAB ROW INDEX
      var index = Number(event.target.parentNode.parentNode.rowIndex);
      //console.log(index);

    //CONFIRM SETUP
    alertify.set({ 
      labels: {
        ok:"Yes",
        cancel:"No"
      },
      buttonReverse:true

    });

    //CONFIRM
    alertify.confirm("Are you sure you want to delete the item?", function(e){

      //YES
      if(e){

        //AJAX
        $.ajax({
           type: "POST",
           url: url,
           success: function(res){
              //RESET FORM              
              ResetForm(); 
              
              //DELETE ROW
              $("#List")[0].deleteRow(index);
               
              //ALERT
              alertify.success(res);
           }

        //END AJAX
        });
      }

    //END CONFIRM       
    });
      

       

  //END CLICK
  }); 
   //

*/




/*********TABS************/
  //1ST BTN ACTIVE
  $("#btn1").addClass('active');

  //TABS BUTTONS
  $("#buttons").on('click',btnClick);
  
  //TABS BUTTON CLICK      
  function btnClick(e){
        
    e.preventDefault();

      //console.log(e);

    if(e.target.tagName == "A"){
      var hash = e.target.hash;
      var num = hash.substr(-1);
      var btns = e.currentTarget.children;
      
      $.each(btns,function(key,val){
          var tab = "#tab"+String(key+1);
            $(tab).css('zIndex',0);
          var btn = $(btns[key]);
              btn[0].className = "";
      });


      $(hash).css('zIndex',1);
      $("#btn"+num).addClass('active');
    
    //END IF
    }


  //END BUTTONS CLICK
  }

/*********PANEL VARS************/

//VARS
var AdminMenuClosed = false;
var SideBarClosed = false;
var contentDiv = $("#content");


/*********ADMIN MENU************/

//CLICK SIDEBAR CLOSE BUTTON
$("#MenuLogo").on('click',function(event){

  //PREVENT DEFAULT
  event.preventDefault();


  //CLOSE
  if(AdminMenuClosed == false){

    $("#AdminMenu").css('margin-left','-300px');
    $("#shell").css('margin-left','10px');
    $("#PageBarGutter").css('width','10px');

    AdminMenuClosed = true;

    if(SideBarClosed == true && AdminMenuClosed == true){
      contentDiv.width(String(contentDiv.width() + 290) + "px");
    
    } else if(SideBarClosed == false && AdminMenuClosed == true){
      contentDiv.width(String(contentDiv.width() + 290) + "px");

    } else {

    }


  //OPEN
  } else {

    $("#AdminMenu").css('margin-left','0px');
    $("#shell").css('margin-left','300px');
    $("#PageBarGutter").css('width','300px');

    AdminMenuClosed = false;


    if(SideBarClosed == true && AdminMenuClosed == false){
      contentDiv.width(String(contentDiv.width() - 290) + "px");
    
    } else if(SideBarClosed == false && AdminMenuClosed == false){
      contentDiv.width(String(contentDiv.width() - 290) + "px");

    } else {

    }

  }

});

/*********SIDEBAR************/
if(SideBarClosed == false){
  contentDiv.width(String(contentDiv.width() - 332) + "px");
}


//CLICK SIDEBAR CLOSE BUTTON
$("#SidebarCloseBtn").on('click',function(event){

  //PREVENT DEFAULT
  event.preventDefault();

  //CLOSED
  if(SideBarClosed == false){


    $("#Sidebar").css('margin-right','-312px');
    $(this).html(" <<br />< ");
    contentDiv.width(String(contentDiv.width() + 312)+"px");
    SideBarClosed = true;

  //OPEN
  } else {
  
    $("#Sidebar").css('margin-right','0px');
    $(this).html(" ><br />> ");
    contentDiv.width(String(contentDiv.width() - 312) + "px");


    SideBarClosed = false;

  }

});


//END DOC
});