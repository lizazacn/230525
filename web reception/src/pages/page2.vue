<template>
  <div style="width: 300px">
    <el-form
        ref="formRef"
        :model="classValidateForm"
        label-width="100px"
        class="demo-ruleForm"
    >
      <el-form-item
          label="班级"
          prop="ClassNumber"
          :rules="[
        { required: true, message: 'ClassNumber is required' },
      ]"
      >
        <el-input
            v-model.number="classValidateForm.ClassNumber"
            type="text"
            autocomplete="off"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">Submit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance } from 'element-plus'
import {useRouter} from 'vue-router'
import {ElLoading} from "element-plus";

const formRef = ref<FormInstance>()
const router = useRouter()

const classValidateForm = reactive({
  ClassNumber: '',
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      // let time =6
      // setInterval(function (){
      //
      // },1000)
      // const loading = ElLoading.service({
      //   lock: true,
      //   text: 'Loading',
      //   background: 'rgba(0, 0, 0, 0.7)',
      // })

      router.push({name:'page3'})
    } else {
      console.log('error submit!')
      return false
    }
  })
}

</script>
