<template>
  <div style="width: 300px">
    <el-form
        ref="formRef"
        :model="classValidateForm"
        label-width="100px"
        class="demo-ruleForm"
    >
      <el-form-item
          label="性别"
          prop="sex"
          :rules="[
        { required: true, message: 'sex is required' },
      ]"
      >
        <el-input
            v-model.number="classValidateForm.sex"
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

const formRef = ref<FormInstance>()
const router = useRouter()

const classValidateForm = reactive({
  sex: '',
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      console.log('submit!')
      router.push({name:'page2'})
    } else {
      console.log('error submit!')
      return false
    }
  })
}

</script>
