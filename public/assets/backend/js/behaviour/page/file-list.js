function ownerTypeChange(a){}
function applySelected(){
	var s = getSelectedFiles();
	if(!s) return false;
	if(client=='xheditor'&&window.callback){
		window.callback('!'+s.files.join(' '));
		return false;
	}
	if(callback){
		if(typeof(target[callback])=='function'){
			target[callback](s.files,s.infos);
		}
	}else{
		if(insertTo){
			$(insertTo).val(s.files.join(','));
			return false;
		}
        App.message({title:App.i18n.SYS_INFO,text:App.i18n.NO_CALLBACK_NAME,type:'error'});
        return false;
	}
    return false;
}
function getSelectedFiles(){
	var files=[],infos=[];
	$("input.check-table:checked").each(function(){
		var info=$(this).data('info');
		//{"id":2026,"owner_type":"user","owner_id":1,"name":"未标题-1.jpg","save_name":"102568426320429056.jpg","save_path":"public/upload/b2c_store_goods/0/102568426320429056.jpg","view_url":"/public/upload/b2c_store_goods/0/102568426320429056.jpg","ext":".jpg","mime":"image/jpeg","type":"image","size":26341,"width":388,"height":388,"dpi":0,"md5":"8541a261e493e3b15f982c361d3f97cc","storer_name":"local","storer_id":"","created":1596896034,"updated":0,"project":"","table_id":"0","table_name":"official_b2c_store_goods","field_name":"content","sort":0,"status":0,"category_id":0,"used_times":0}
		files.push(info.view_url);
		infos.push(info);
	});
	if(files.length<1){
		App.message({title:App.i18n.SYS_INFO,text:App.i18n.PLEASE_SELECT,type:'error'});
		return false;
	}
	return {files:files,infos:infos};
}
$(function(){
	$('#timerange').on('focus',function(){
		if($(this).data('attached')) return false;
		$(this).data('attached',true);
		App.daterangepicker('#timerange',{
			showShortcuts: true,
			shortcuts: {
				'prev-days': [1,3,5,7],
				'next-days': [3,5,7],
				'prev' : ['week','month'],
				'next' : ['week','month']
			}
		});
	});
	function submitSearch(e){
		e.preventDefault();
		var data=$('#search-form').serializeArray();
		data.push({name:'partial',value:1});
		loadList($('#search-form').attr('action'),data);
	}
	if(dialogMode) {
		$('#search-form').on('submit',submitSearch);
		$('#timerange,#type,#table,#ownerId,#used').on('change', submitSearch);
	}else{
		$('#timerange,#type,#table,#ownerId,#used').on('change', function(){
			$('#search-form').submit();
		});
	}
	function initUploadButton(){
		App.uploadPreviewer("#input-file-upload", {url:uploadURL}, function(r){
			if(r.Code==1) {
				loadList(listURL,{partial:1});
			}
		});
		$('#checkedAll,input[type=checkbox][name="id[]"]:checked').prop('checked',false);
		App.attachCheckedAll('#checkedAll','input[type=checkbox][name="id[]"]');
	}
	function loadList(url,args){
		$.get(url,args,function(r){
			$('#table-container').html(r);
			initTable();
			$('#table-container .pagination a').on('click',function(e){
				e.preventDefault();
				var url=$(this).attr('href');
				loadList(url,{});
			});
		},'html');
	}
	function initTable(){
		$('#table-container thead').data('sort-trigger',function(){
			var thead=$('#table-container thead');
			var url=thead.data('sort-url');
			loadList(url,{partial:1});
		});
		App.tableSorting('#table-container');
		App.float('#tbody-content img.previewable', null, null, '8-6');
	}
	initUploadButton();
	initTable();
});