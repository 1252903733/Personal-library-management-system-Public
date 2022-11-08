<template>

<!--  <el-col :span="24">
    <el-upload
        ref="upload"
        class="upload-demo"
        :limit="1"
        multiple
        :http-request="uploa    dfile"
        list-type="picture"
    >
      <template #trigger>
        <el-button type="primary">请选择修改后的课程文档(可不选择修改)</el-button>
      </template>
      &lt;!&ndash;                <template #tip>
                          <div class="el-upload__tip text-red">
                              上传不能多个文件  上传失败
                          </div>
                      </template>&ndash;&gt;
    </el-upload>
  </el-col>-->
  <!--    查询学生信息-->
  <el-row :gutter="20">
    <el-col :span="3">
      <el-form-item label="图书编号">
        <el-input v-model="querybook.Id"></el-input>
      </el-form-item>
    </el-col>
    <el-col :span="3">
      <el-form-item label="书籍作者">
        <el-input v-model="querybook.Author"></el-input>
      </el-form-item>
    </el-col>
    <el-col :span="3">
      <el-form-item label="书籍名称">
        <el-input v-model="querybook.Name"></el-input>
      </el-form-item>
    </el-col>
    <el-col :span="3">
      <el-form-item label="书籍内容">
        <el-input v-model="querybook.Content"></el-input>
      </el-form-item>
    </el-col>
    <el-col :span="3">
        <el-button type="success" size="large" @click="addbookfunction()">新增图书信息</el-button>
    </el-col>
    <el-col :span="6">
      <div style="font-size: 25px;color:black">当前搜索书籍数量:{{page.total}}</div>
    </el-col>
  </el-row>

  <div style="background-color:cornflowerblue;">  <el-pagination
      v-model:currentPage="page.currentpage"
      v-model:page-size="page.pagesize"
      :page-sizes="[8,7,6,5,4,3,2,1]"
      background
      layout="total, sizes, prev, pager, next, jumper"
      :total="page.total"
      @size-change="handleSizeChange(val)"
      @current-change="handleCurrentChange(val)"
  /></div>
  <el-table :header-cell-style="{background:'rgb(34,204,254)' ,color:'white',fontSize:'20px'}"   :data="book.slice((page.currentpage-1)*page.pagesize,page.currentpage*page.pagesize)" border="true" style="width: 100%;color:black;font-size: 20px" :cell-class-name="tableRowClassName">
    <el-table-column  prop="Id" disabled="" label="图书编号" ></el-table-column>
    <el-table-column  prop="Name"  label="图书名称" ></el-table-column>
    <el-table-column   label="照片" >
      <template #default="scope">
        <el-image style="width: 80px;height: 100px" :src="scope.row.Src"></el-image>
      </template>
    </el-table-column>
    <el-table-column min-width="150%" label="操作" >
      <template #default="scope">
        <el-button type="primary" size="large" @click="revisebookfunction(scope.row)">编辑</el-button>
        <el-button type="danger" size="large" @click="deletebookfunction(scope.row)">删除</el-button>
        <el-button type="success" size="large" @click="browsefunction(scope.row)">在线预览资料</el-button>
        <el-button type="warning" size="large" @click="downloafile(scope.row)">文档下载</el-button>
      </template>
    </el-table-column>
    <el-table-column  prop="Author"  label="书籍作者" ></el-table-column>
    <el-table-column  prop="Content"  label="书籍内容" ></el-table-column>
  </el-table>
  <!--下面是查看班级学生-->
  <el-drawer v-model="showbrowse"  style="padding: 0px 0px;border: none" :show-close="false" :direction="direction" size="70%">
     <div id="shwofile"></div>
  </el-drawer>
  <!--下面是查看班级学生-->
  <el-drawer v-model="showpdf"  style="padding: 0px 0px;border: none" :show-close="false" :direction="direction" size="70%">
    <iframe :src="pdfurl" width="1400px" height="1000px"></iframe>
  </el-drawer>







  <el-dialog v-model="addbookvisiable" title="新增图书信息">
    <el-form v-model="addbook">
      <el-row :gutter="20">
        <el-col :span="24">
          <el-upload
              ref="upload"
              class="upload-demo"
              :limit="1"
               accept=".docx,.pdf"
              :on-exceed="handleExceed"
              :http-request="uploadfile"
              list-type="picture"
          >
            <template #trigger>
              <el-button type="primary">请选择图书文档资料(仅仅支持pdf,docx格式)</el-button>
            </template>
          </el-upload>

          <el-upload
              ref="upload"
              class="upload-demo"
              accept=".jpg,.png"
              :limit="1"
              multiple
              :on-exceed="handleExceed"
              :http-request="uploadfile1"
              list-type="picture"
          >
            <template #trigger>
              <el-button type="success">请选择图书照片展示(仅仅支持jpg,png格式)</el-button>
            </template>
          </el-upload>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书编号">
            <el-input v-model="addbook.Id" size="large"  @input="handleEdit"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
        <el-form-item label="图书内容">
          <el-input v-model="addbook.Content"></el-input>
        </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书名称">
            <el-input v-model="addbook.Name" size="large"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书作者">
            <el-input v-model="addbook.Author" size="large"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-button type="warning" @click="confirmaddbooktrigger()">确认添加</el-button>
        </el-col>
      </el-row>
    </el-form>
  </el-dialog>




  <el-dialog v-model="revisebookvisiable" title="新增图书信息">
    <el-form v-model="revisebook">
      <el-row :gutter="20">
        <el-col :span="24">
          <el-upload
              ref="upload"
              class="upload-demo"
              :limit="1"
              accept=".docx,.pdf"
              :on-exceed="handleExceed"
              :http-request="reviseuploadfile"
              list-type="picture"
          >
            <template #trigger>
              <el-button type="primary">请选择修改后的图书文档资料(仅仅支持pdf,docx格式)可选择不修改</el-button>
            </template>
          </el-upload>

          <el-upload
              ref="upload"
              class="upload-demo"
              accept=".jpg,.png"
              :limit="1"
              multiple
              :on-exceed="handleExceed"
              :http-request="reviseuploadfile1"
              list-type="picture"
          >
            <template #trigger>
              <el-button type="success">请选择修改后的图书照片展示(仅仅支持jpg,png格式)可选择不修改</el-button>
            </template>
          </el-upload>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书编号">
            <el-input v-model="revisebook.Id" disabled size="large"  @input="handleEdit"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书内容">
            <el-input v-model="revisebook.Content"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书名称">
            <el-input v-model="revisebook.Name" size="large"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="图书作者">
            <el-input v-model="revisebook.Author" size="large"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-button type="warning" @click="confirmrevisebooktrigger()">确认修改</el-button>
        </el-col>
      </el-row>
    </el-form>
  </el-dialog>



</template>




<script>
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {postfile,get,getfile,post} from "@/util/service";
import { reactive, onMounted, ref, toRefs,watch } from 'vue'
import { ElMessageBox,ElMessage} from 'element-plus'
import store from "@/store/indexstore";

/*
import pdf from "vue-pdf"
*/
export default {
 /* components:{
    /!*pdf*!/
  },*/
  setup() {
    const router=useRouter()
    const shwofile=ref(null)
    const docx=require("docx-preview")
    window.JSZip=require("jszip")
    const data = reactive({
      addbookvisiable:false,
      revisebookvisiable:false,
      srcfile:"",
      docfile:"",
      showbrowse:false,
      showpdf:false,
      pdfurl:"",
      direction:"rtl",
      url:'',
      bookstorage:"",
      page:{
        currentpage:1,//当前页面
        pagesize:5,//每页的数据个数
        total:10//数据总数
      },
      book:"",
      bookres:"",
      file:"",
      querybook:{
        Id:"",
        Name:"",
        Content:"",
        Author:"",
      },
      addbook:{
        Id:"",
        Name: "",
        Author: "",
        Content: "",
      },
      revisebook:{
        Id:"",
        Name: "",
        Author: "",
        Content: "",
      },
      revisesrcfile:"",
      revisedocfile:"",

    })

    onMounted(async () => {//初始化函数
      await get("/getbookdata").then(res => {
        data.bookstorage=JSON.parse(JSON.stringify(res))
        data.bookres =JSON.parse(JSON.stringify(res))
        console.log("bookstorage1数据为",data.bookstorage)
        console.log("bookres1数据为",data.bookres)

//获取学生信息 然后进行显示
      })
         for (const key in data.bookres) {
          const option = {
          Src: data.bookres[key]['Src'],
          Id1: data.bookres[key]['Id'] //必修标记为Name不然gin 框架接受不到
        }
        getfile("/getfile", option).then(res => {
          data.bookres[key]['Src']=window.URL.createObjectURL(res)
           })
           console.log("bookstorage数据为",data.bookstorage)
           console.log("bookres数据为",data.bookres)
        data.book = data.bookres
        data.page.total = data.book.length
         }



    })
/*

    const uploadfile=async (param) => {
      let file = param.file;//上传之后获取这个上传的文件信息s
      data.file = file//将这个信息赋值给data.coursedocument*!/

      //这个是获取图片路径进行显示  实际就是将后端的图片转换成二进制流  再转换成图片的过程

    }
*/


//监听操作  只要有查询变化就触发操作
    watch(()=>[data.querybook],(newvalue,oldvalue)=>{
          data.book=[]
          let i=0
          for(const key in data.bookres)
          {
            if(isincludes(data.bookres[key],data.querybook)===1)
            {
              data.book[i]=data.bookres[key]
              i=i+1//从0开始  最后有i个数据
            }
          }
          data.page.total=i
        },
        {
          deep:true//开启深度监听
        }
    )
    onMounted( async () => {
      if (store.state.CurrentStatusInformation !== "" && store.state.CurrentStatusInformation !== null) {
        alert(store.state.CurrentStatusInformation)
        localStorage.setItem('currentstatusinformation', "")//当前状态信息
      }
      })



    const browsefunction=(param)=>{
      const option = {
        Id1: param.Id,
        DocSrc: param.DocSrc//必修标记为Name不然gin 框架接受不到
      }
      getfile("/getdoc", option).then(res => {
        if(param.DocSrc==="docx") {
          docx.renderAsync(res, document.getElementById("shwofile"))
        }
        else
        {
          data.pdfurl=  "http://localhost:8019/getdow?Id1="+param.Id+"&DocSrc=pdf"
          console.log( data.pdfurl)

/*
     data.pdfurl ="http://localhost:8019/getdow?Id1=1&DocSrc=pdf"
*/
        }
      })
      if(param.DocSrc==="docx")
      {
        console.log("是docx")
        data.showbrowse=true
      }
      else
      {
        console.log("是pdf")
        data.showpdf=true
      }
    }

    const handleEdit = (e) => {
      //正则表达式输入非0开头的数字1-9位
      let num = e.toString()
      var reg = /^((?!0)\d{1,9})$/
      if (!num.match(reg)) {
      data.addbook.Id= data.addbook.Id.slice(0, data.addbook.Id.length-1)
      } else {
        data.addbook.Id=num
      }
    }

    const downloafile = (param) => {
      const option = {
        Id1: param.Id,
        DocSrc: param.DocSrc//必修标记为Name1不然gin 框架接受不到
      }
      getfile("/getdow", option).then(res => {
        //res是一个bog对象 解析这个对象生成图篇数据
        let url=window.URL.createObjectURL(res)
        const aLink = document.createElement('a')
        aLink.href = url
        // 2.直接使用自定义文件名,设置下载文件名称
        aLink.setAttribute('download', param.Name+"."+param.DocSrc )
        document.body.appendChild(aLink)
        // 模拟点击下载
        aLink.click()
        // 移除改下载标签
        document.body.removeChild(aLink);
      })
    }

const addbookfunction= () => {
  data.addbookvisiable=true
}

const revisebookfunction = (param) => {
      data.revisebook.Id=param.Id
      data.revisebook.Name=param.Name
      data.revisebook.Author=param.Author
      data.revisebook.Content=param.Content
      data.revisebook.Src=param.Src
      data.revisebook.DocSrc=param.DocSrc
      console.log("获取docsoc",param.DocSrc)
      data.revisebookvisiable=true

}



    const confirmaddbooktrigger = () => {


      for(const key in data.addbook)
      {
       if(data.addbook[key]==="")
        {
          ElMessage.error("输入信息不能为空")
          return
        }
      }

      if(data.srcfile==="")
      {
        ElMessage.error("图书照片没有上传")
        return
      }



      if(data.docfile==="")
      {
        ElMessage.error("图书资料没有上传")
        return
      }


      for(const key in data.bookres)
      {
        if(String(data.bookres[key]['Name'])===String(data.addbook.Name))
        {
          ElMessage.error("该图书名称已经存在")
          return
        }
        if(String(data.bookres[key]['Id'])===String(data.addbook.Id)) {
          console.log("遍历图书名称为:"+data.bookres[key]['Id'])
          console.log("添加图书名称为:"+data.addbook.Id)
          ElMessage.error("该图书编号已经存在")
          return
        }
      }



      var formdata =new FormData()
      formdata.append("Id1",data.addbook.Id)
      formdata.append("Name1",data.addbook.Name)//go后端接受name会冲突
      formdata.append("Content",data.addbook.Content)
      formdata.append("Author",data.addbook.Author)
      formdata.append("Src",data.srcfile.name.substring(data.srcfile.name.lastIndexOf(".")+1))
      formdata.append("DocSrc",data.docfile.name.substring(data.docfile.name.lastIndexOf(".")+1))
      formdata.append("srcfile",data.srcfile)
      formdata.append("docsrcfile",data.docfile)
      post("/add", formdata).then(res => {
        console.log("gst返回数据为",res)
         if(res==="success"){
           localStorage.setItem('currentstatusinformation',"添加信息成功")
           router.go(0)
         }
         else {
         localStorage.setItem('currentstatusinformation',"添加图书信息失败")
           router.go(0)

         }
      })

    }
    const deletebookfunction = (param) => {
       ElMessageBox.confirm(
          '确定修改吗?',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      ).then((action)=>{
        let formdata=new FormData()
         formdata.append("Id1",param.Id)
         post("/delbook", formdata).then(res => {
           if(res==="success"){
             localStorage.setItem('currentstatusinformation',"修改图书信息成功")
             router.go(0)
           }
           else {
             localStorage.setItem('currentstatusinformation',"修改图书信息失败")
             router.go(0)
           }
         })
      })
    }

const confirmrevisebooktrigger = () => {
  ElMessageBox.confirm(
      '确定修改吗?',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then((action)=>{
    if(action=="confirm")//说明点击确定
    {
      var formdata =new FormData()

      formdata.append("Id1",data.revisebook.Id)
      formdata.append("Name1",data.revisebook.Name)//go后端接受name会冲突
      formdata.append("Content",data.revisebook.Content)
      formdata.append("Author",data.revisebook.Author)

      if(data.revisesrcfile!=="")
      {
        formdata.append("Src",data.revisesrcfile.name.substring(data.revisesrcfile.name.lastIndexOf(".")+1))
        formdata.append("srcfile",data.revisesrcfile)

      }
      else {
        for(const key in data.bookstorage)
        {
          if(String(data.bookstorage[key]['Id'])===String(data.revisebook.Id))
          {
            formdata.append("Src",data.bookstorage[key]['Src'])
          }
        }
      }

      if(data.revisedocfile!=="")
      {
        formdata.append("DocSrc",data.revisedocfile.name.substring(data.revisedocfile.name.lastIndexOf(".")+1))
        formdata.append("docsrcfile",data.revisedocfile)
      }else
      {
        formdata.append("DocSrc",data.revisebook.DocSrc)
      }

      post("/revise", formdata).then(res => {
        if(res==="success"){
          localStorage.setItem('currentstatusinformation',"修改图书信息成功")
          router.go(0)
        }
        else {
          localStorage.setItem('currentstatusinformation',"修改图书信息失败")
          router.go(0)

        }
      })

    }
  })
}








//根据
    const isincludes=(data,data1)=>
    {
      for(const key in data)
      {
        if(key!=="Src"&&key!=="DocSrc"&&!String(data[key]).includes(String(data1[key])))
        {
          return 0
        }
      }
      return 1
    }


    const uploadfile = (param) => {
      data.docfile=param.file
    }



    const uploadfile1 = (param) => {
      data.srcfile=param.file
    }

    const reviseuploadfile = (param) => {
      data.revisedocfile=param.file



    }


    const reviseuploadfile1 = (param) => {
      data.revisesrcfile=param.file

    }




    const handleSizeChange=(val)=>{
      console.log(data.page.pagesize)
    }

    const handleCurrentChange=(val)=>
    {
      console.log(data.page.currantpage)
    }


    return{
      ...toRefs(data),
      confirmrevisebooktrigger,
      deletebookfunction,
      revisebookfunction,
      handleSizeChange,
      handleCurrentChange,
      downloafile,
      addbookfunction,
      uploadfile1,
      uploadfile,
      reviseuploadfile,
      reviseuploadfile1,
      confirmaddbooktrigger,
      handleEdit,
     browsefunction
    }
    //


  }
}



</script>